package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

type CallResult struct {
	OK    bool        `json:"ok"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type CallFunc func(params json.RawMessage) *CallResult

type FuncInfo struct {
	Name   string   `json:"name"`
	Params []string `json:"params,omitempty"`
	Desc   string   `json:"desc,omitempty"`
}

type Registry struct {
	mu    sync.RWMutex
	funcs map[string]CallFunc
	infos map[string]*FuncInfo
}

var globalRegistry = &Registry{
	funcs: make(map[string]CallFunc),
	infos: make(map[string]*FuncInfo),
}

func (r *Registry) Register(name string, paramNames []string, desc string, fn CallFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.funcs[name] = fn
	r.infos[name] = &FuncInfo{Name: name, Params: paramNames, Desc: desc}
}

func (r *Registry) Call(name string, jsonParams string) (result *CallResult) {
	defer func() {
		if rec := recover(); rec != nil {
			result = &CallResult{Error: fmt.Sprintf("panic: %v", rec)}
		}
	}()

	r.mu.RLock()
	fn, ok := r.funcs[name]
	r.mu.RUnlock()

	if !ok {
		return &CallResult{Error: "function not found: " + name}
	}

	return fn(json.RawMessage(jsonParams))
}

func (r *Registry) List() []*FuncInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*FuncInfo, 0, len(r.infos))
	for _, info := range r.infos {
		list = append(list, info)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

func okResult(data interface{}) *CallResult {
	return &CallResult{OK: true, Data: data}
}

func errResult(msg string) *CallResult {
	return &CallResult{Error: msg}
}

func errResultf(format string, args ...interface{}) *CallResult {
	return &CallResult{Error: fmt.Sprintf(format, args...)}
}
