package utils

import (
	"net/url"
	"strings"
)

// Post数据类 HTTP POST 请求参数构建器。
// 用于构建 key=value 格式的 POST 表单数据或协议头数据。
// 内部维护 keys 和 values 两个平行的字符串数组，保持插入顺序。
type Post数据类 struct {
	keys   []string
	values []string
}

// T添加 向 Post 数据中添加一个键值对。
//
// 参数:
//   - key: 参数名
//   - value: 参数值
//   - 转码: true 时对 value 进行 URL 编码，false 时保留原值
func (p *Post数据类) T添加(key, value string, 转码 bool) {
	p.keys = append(p.keys, key)
	if 转码 {
		p.values = append(p.values, url.QueryEscape(value))
	} else {
		p.values = append(p.values, value)
	}
}

// T添加_批量 从 "key1=value1&key2=value2" 格式的文本批量添加键值对。
//
// 参数:
//   - 文本: 以 & 分隔、= 连接的键值对文本
//   - 转码: true 时对 value 进行 URL 编码
func (p *Post数据类) T添加_批量(文本 string, 转码 bool) {
	arr := strings.Split(文本, "&")
	for _, item := range arr {
		keyValue := strings.Split(item, "=")
		key := keyValue[0]
		value := keyValue[1]
		p.T添加(key, value, 转码)
	}
}

// Q取值 根据键名获取对应的值。
//
// 参数:
//   - key: 要查找的参数名
//
// 返回:
//   - string: 对应的参数值；键名不存在时返回空串
func (p *Post数据类) Q取值(key string) string {
	for i, k := range p.keys {
		if k == key {
			return p.values[i]
		}
	}
	return ""
}

// Z置值 设置指定键名的值。如果键名已存在则更新，不存在则添加。
//
// 参数:
//   - key: 参数名
//   - value: 新的参数值
func (p *Post数据类) Z置值(key, value string) {
	for i, k := range p.keys {
		if k == key {
			p.values[i] = value
			return
		}
	}
	p.T添加(key, value, false)
}

// H获取Post数据 将所有键值对拼接为 POST 表单格式的字符串。
// 格式：key1=value1&key2=value2
//
// 参数:
//   - 是否URL编码: true 时对所有 value 进行 URL 编码
//
// 返回:
//   - string: 拼接后的 POST 数据字符串
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

// H获取协议头数据 将所有键值对拼接为 HTTP 协议头格式的字符串。
// 格式：key1: value1\r\nkey2: value2
//
// 参数:
//   - 是否URL编码: true 时对所有 value 进行 URL 编码
//
// 返回:
//   - string: 拼接后的协议头字符串
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

// H获取Key数组 返回所有键名的数组副本。
//
// 返回:
//   - []string: 键名数组
func (p *Post数据类) H获取Key数组() []string {
	return p.keys
}

// H获取Value数组 返回所有值的数组副本。
//
// 返回:
//   - []string: 值数组
func (p *Post数据类) H获取Value数组() []string {
	return p.values
}

// Q清空 清除所有键值对，重置为空状态。
func (p *Post数据类) Q清空() {
	p.keys = nil
	p.values = nil
}

// S删除 删除指定键名的键值对。
// 如果键名不存在，不做任何操作。
//
// 参数:
//   - key: 要删除的参数名
func (p *Post数据类) S删除(key string) {
	for i := 0; i < len(p.keys); i++ {
		if p.keys[i] == key {
			p.keys = append(p.keys[:i], p.keys[i+1:]...)
			p.values = append(p.values[:i], p.values[i+1:]...)
			return
		}
	}
}

// H获取JSON文本 将所有键值对拼接为 JSON 对象格式的字符串。
// 格式：{"key1":"value1","key2":"value2"}
// 注意：值中的特殊字符（如双引号）未做转义处理，复杂场景请使用 encoding/json。
//
// 返回:
//   - string: JSON 格式的字符串
func (p *Post数据类) H获取JSON文本() string {
	str := "{"
	for i := 0; i < len(p.keys); i++ {
		str += "\"" + p.keys[i] + "\":\"" + p.values[i] + "\","
	}
	str = str[:len(str)-1] + "}"
	return str
}
