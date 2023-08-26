package utils

import (
	"net/url"
	"strings"
)

type Post数据类 struct {
	keys   []string
	values []string
}

func (p *Post数据类) T添加(key, value string, 转码 bool) {
	p.keys = append(p.keys, key)
	if 转码 {
		p.values = append(p.values, url.QueryEscape(value))
	} else {
		p.values = append(p.values, value)
	}
}

func (p *Post数据类) T添加_批量(文本 string, 转码 bool) {
	arr := strings.Split(文本, "&")
	for _, item := range arr {
		keyValue := strings.Split(item, "=")
		key := keyValue[0]
		value := keyValue[1]
		p.T添加(key, value, 转码)
	}
}

func (p *Post数据类) Q取值(key string) string {
	for i, k := range p.keys {
		if k == key {
			return p.values[i]
		}
	}
	return ""
}

func (p *Post数据类) Z置值(key, value string) {
	for i, k := range p.keys {
		if k == key {
			p.values[i] = value
			return
		}
	}
	p.T添加(key, value, false)
}

func (p *Post数据类) H获取Post数据(是否URL编码 bool) string {
	str := ""
	for i := 0; i < len(p.keys); i++ {
		if 是否URL编码 {
			str += p.keys[i] + "=" + url.QueryEscape(p.values[i]) + "&"
		} else {
			str += p.keys[i] + "=" + p.values[i] + "&"
		}
	}
	str = str[:len(str)-1]
	return str
}

func (p *Post数据类) H获取协议头数据(是否URL编码 bool) string {
	str := ""
	for i := 0; i < len(p.keys); i++ {
		if 是否URL编码 {
			str += p.keys[i] + ": " + url.QueryEscape(p.values[i]) + "\r\n"
		} else {
			str += p.keys[i] + ": " + p.values[i] + "\r\n"
		}
	}
	str = str[:len(str)-2]
	return str
}

func (p *Post数据类) H获取Key数组() []string {
	return p.keys
}

func (p *Post数据类) H获取Value数组() []string {
	return p.values
}

func (p *Post数据类) Q清空() {
	p.keys = nil
	p.values = nil
}

func (p *Post数据类) S删除(key string) {
	for i := 0; i < len(p.keys); i++ {
		if p.keys[i] == key {
			p.keys = append(p.keys[:i], p.keys[i+1:]...)
			p.values = append(p.values[:i], p.values[i+1:]...)
			return
		}
	}
}

func (p *Post数据类) H获取JSON文本() string {
	str := "{"
	for i := 0; i < len(p.keys); i++ {
		str += "\"" + p.keys[i] + "\":\"" + p.values[i] + "\","
	}
	str = str[:len(str)-1] + "}"
	return str
}
