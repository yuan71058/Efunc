// PostData 类 - HTTP POST 请求参数构建器
// 用于构建 key=value 格式的 POST 表单数据或协议头数据。
// 内部维护 keys 和 values 两个平行的字符串数组，保持插入顺序。
package utils

import (
	"net/url"
	"strings"
)

// PostData class HTTP POST 请求参数构建器。
type PostData struct {
	keys   []string
	values []string
}

// Add 向 Post 数据中添加一个键值对。
//
// 参数:
//   - key: 参数名
//   - value: 参数值
//   - encode: true 时对 value 进行 URL 编码，false 时保留原值
func (p *PostData) Add(key, value string, encode bool) {
	p.keys = append(p.keys, key)
	if encode {
		p.values = append(p.values, url.QueryEscape(value))
	} else {
		p.values = append(p.values, value)
	}
}

// AddBatch 从 "key1=value1&key2=value2" 格式的文本批量添加键值对。
//
// 参数:
//   - text: 以 & 分隔、= 连接的键值对文本
//   - encode: true 时对 value 进行 URL 编码
func (p *PostData) AddBatch(text string, encode bool) {
	arr := strings.Split(text, "&")
	for _, item := range arr {
		keyValue := strings.Split(item, "=")
		key := keyValue[0]
		value := keyValue[1]
		p.Add(key, value, encode)
	}
}

// Get 根据键名获取对应的值。
//
// 参数:
//   - key: 要查找的参数名
//
// 返回:
//   - string: 对应的参数值；键名不存在时返回空串
func (p *PostData) Get(key string) string {
	for i, k := range p.keys {
		if k == key {
			return p.values[i]
		}
	}
	return ""
}

// Set 设置指定键名的值。如果键名已存在则更新，不存在则添加。
//
// 参数:
//   - key: 参数名
//   - value: 新的参数值
func (p *PostData) Set(key, value string) {
	for i, k := range p.keys {
		if k == key {
			p.values[i] = value
			return
		}
	}
	p.Add(key, value, false)
}

// GetPostData 将所有键值对拼接为 POST 表单格式的字符串。
// 格式：key1=value1&key2=value2
//
// 参数:
//   - urlEncode: true 时对所有 value 进行 URL 编码
//
// 返回:
//   - string: 拼接后的 POST 数据字符串
func (p *PostData) GetPostData(urlEncode bool) string {
	str := ""
	for i := 0; i < len(p.keys); i++ {
		if urlEncode {
			str += p.keys[i] + "=" + url.QueryEscape(p.values[i]) + "&"
		} else {
			str += p.keys[i] + "=" + p.values[i] + "&"
		}
	}
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

// GetHeaderData 将所有键值对拼接为 HTTP 协议头格式的字符串。
// 格式：key1: value1\r\nkey2: value2
//
// 参数:
//   - urlEncode: true 时对所有 value 进行 URL 编码
//
// 返回:
//   - string: 拼接后的协议头字符串
func (p *PostData) GetHeaderData(urlEncode bool) string {
	str := ""
	for i := 0; i < len(p.keys); i++ {
		if urlEncode {
			str += p.keys[i] + ": " + url.QueryEscape(p.values[i]) + "\r\n"
		} else {
			str += p.keys[i] + ": " + p.values[i] + "\r\n"
		}
	}
	if len(str) >= 2 {
		str = str[:len(str)-2]
	}
	return str
}

// GetKeys 返回所有键名的切片。
//
// 返回:
//   - []string: 键名切片
func (p *PostData) GetKeys() []string {
	return p.keys
}

// GetValues 返回所有值的切片。
//
// 返回:
//   - []string: 值切片
func (p *PostData) GetValues() []string {
	return p.values
}

// Clear 清除所有键值对，重置为空状态。
func (p *PostData) Clear() {
	p.keys = nil
	p.values = nil
}

// Remove 删除指定键名的键值对。
// 如果键名不存在，不做任何操作。
//
// 参数:
//   - key: 要删除的参数名
func (p *PostData) Remove(key string) {
	for i := 0; i < len(p.keys); i++ {
		if p.keys[i] == key {
			p.keys = append(p.keys[:i], p.keys[i+1:]...)
			p.values = append(p.values[:i], p.values[i+1:]...)
			return
		}
	}
}

// GetJSON 将所有键值对拼接为 JSON 对象格式的字符串。
// 格式：{"key1":"value1","key2":"value2"}
// 注意：值中的特殊字符（如双引号）未做转义处理，复杂场景请使用 encoding/json。
//
// 返回:
//   - string: JSON 格式的字符串
func (p *PostData) GetJSON() string {
	str := "{"
	for i := 0; i < len(p.keys); i++ {
		str += "\"" + p.keys[i] + "\":\"" + p.values[i] + "\","
	}
	if len(str) > 1 {
		str = str[:len(str)-1] + "}"
	}
	return str
}