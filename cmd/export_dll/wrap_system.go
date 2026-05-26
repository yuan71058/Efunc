package main

import (
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func registerSystemFunctions() {
	r := globalRegistry

	r.Register("X系统_取CPU核心数", nil, "获取CPU逻辑核心数",
		func(p json.RawMessage) *CallResult {
			result, err := utils.X系统_取CPU核心数()
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取CPU物理核心数", nil, "获取CPU物理核心数",
		func(p json.RawMessage) *CallResult {
			result, err := utils.X系统_取CPU物理核心数()
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取总CPU使用率", []string{"间隔"}, "获取总CPU使用率",
		func(p json.RawMessage) *CallResult {
			var v struct {
				间隔 int `json:"间隔"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.X系统_取总CPU使用率(v.间隔)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取内存信息", nil, "获取内存使用情况",
		func(p json.RawMessage) *CallResult {
			result, err := utils.X系统_取内存信息()
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取主机信息", nil, "获取主机信息",
		func(p json.RawMessage) *CallResult {
			result, err := utils.X系统_取主机信息()
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取磁盘使用量", []string{"路径"}, "获取磁盘使用量",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"路径"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.X系统_取磁盘使用量(v.路径)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取开机时间", nil, "获取系统开机时长(秒)",
		func(p json.RawMessage) *CallResult {
			result, err := utils.X系统_取开机时间()
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("X系统_取进程名", []string{"pid"}, "根据PID获取进程名",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Pid int32 `json:"pid"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.X系统_取进程名(v.Pid))
		})

	r.Register("X系统_取进程内存占用", []string{"pid"}, "获取进程内存占用",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Pid int32 `json:"pid"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.X系统_取进程内存占用(v.Pid))
		})

	r.Register("X系统_取进程CPU占用", []string{"pid"}, "获取进程CPU占用",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Pid int32 `json:"pid"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.X系统_取进程CPU占用(v.Pid))
		})

	r.Register("X系统_取当前进程ID", nil, "获取当前进程PID",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_取当前进程ID())
		})

	r.Register("X系统_是否64位系统", nil, "判断是否64位系统",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_是否64位系统())
		})

	r.Register("X系统_取系统架构", nil, "获取系统架构",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_取系统架构())
		})

	r.Register("X系统_取操作系统类型", nil, "获取操作系统类型",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_取操作系统类型())
		})

	r.Register("X系统_取逻辑处理器数", nil, "获取逻辑处理器数",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_取逻辑处理器数())
		})

	r.Register("X系统_取Go版本", nil, "获取Go版本",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.X系统_取Go版本())
		})

	r.Register("X线程_延时", []string{"毫秒"}, "线程延时",
		func(p json.RawMessage) *CallResult {
			var v struct {
				毫秒 uint32 `json:"毫秒"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			utils.X线程_延时(v.毫秒)
			return okResult(true)
		})
}

func registerEnvFunctions() {
	r := globalRegistry

	r.Register("K环境_取值", []string{"名称"}, "获取环境变量",
		func(p json.RawMessage) *CallResult {
			var v struct {
				名称 string `json:"名称"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.K环境_取值(v.名称))
		})

	r.Register("K环境_取值带默认值", []string{"名称", "默认值"}, "获取环境变量(带默认值)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				名称   string `json:"名称"`
				默认值 string `json:"默认值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.K环境_取值带默认值(v.名称, v.默认值))
		})

	r.Register("K环境_设置值", []string{"名称", "值"}, "设置环境变量",
		func(p json.RawMessage) *CallResult {
			var v struct {
				名称 string `json:"名称"`
				值   string `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.K环境_设置值(v.名称, v.值); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("K环境_删除值", []string{"名称"}, "删除环境变量",
		func(p json.RawMessage) *CallResult {
			var v struct {
				名称 string `json:"名称"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.K环境_删除值(v.名称); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("K环境_是否存在", []string{"名称"}, "判断环境变量是否存在",
		func(p json.RawMessage) *CallResult {
			var v struct {
				名称 string `json:"名称"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.K环境_是否存在(v.名称))
		})

	r.Register("K环境_取所有", nil, "获取所有环境变量",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.K环境_取所有())
		})

	r.Register("K环境_加载", []string{"文件路径"}, "加载.env文件",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"文件路径"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.K环境_加载(v.路径); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})
}

func registerHTTPFunctions() {
	r := globalRegistry

	r.Register("H客户端_取文本", []string{"网址"}, "HTTP GET获取文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				网址 string `json:"网址"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.H客户端_取文本(v.网址)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("H客户端_下载文件", []string{"网址", "保存路径"}, "下载文件",
		func(p json.RawMessage) *CallResult {
			var v struct {
				网址   string `json:"网址"`
				保存路径 string `json:"保存路径"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.H客户端_下载文件(v.网址, v.保存路径); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})
}
