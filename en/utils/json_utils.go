// JSON 工具
// 基于 tidwall/gjson 和 tidwall/sjson，提供高性能 JSON 路径查询和修改。
// 路径语法：点号分隔，如 "user.name"、"list.0.id"。
package utils

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// JSON_Get 从 JSON 字符串中根据路径获取值。
// 路径语法参考 gjson：点号分隔，如 "user.name"、"list.0.id"。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径，如 "key"、"a.b.c"、"arr.0"
//
// 返回:
//   - string: 路径对应的值（字符串形式）；路径不存在时返回空串
func JSON_Get(jsonText string, path string) string {
	return gjson.Get(jsonText, path).String()
}

// JSON_GetInt 从 JSON 字符串中根据路径获取整数值。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径
//
// 返回:
//   - int64: 路径对应的整数值；路径不存在时返回 0
func JSON_GetInt(jsonText string, path string) int64 {
	return gjson.Get(jsonText, path).Int()
}

// JSON_GetFloat 从 JSON 字符串中根据路径获取浮点数值。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径
//
// 返回:
//   - float64: 路径对应的浮点数值；路径不存在时返回 0
func JSON_GetFloat(jsonText string, path string) float64 {
	return gjson.Get(jsonText, path).Float()
}

// JSON_GetBool 从 JSON 字符串中根据路径获取布尔值。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径
//
// 返回:
//   - bool: 路径对应的布尔值；路径不存在时返回 false
func JSON_GetBool(jsonText string, path string) bool {
	return gjson.Get(jsonText, path).Bool()
}

// JSON_GetArray 从 JSON 字符串中根据路径获取数组。
// 返回 gjson.Result 数组，可通过下标访问每个元素。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径，指向一个 JSON 数组
//
// 返回:
//   - []gjson.Result: 数组元素列表；路径不存在时返回空切片
func JSON_GetArray(jsonText string, path string) []gjson.Result {
	return gjson.Get(jsonText, path).Array()
}

// JSON_GetMap 从 JSON 字符串中根据路径获取对象的键值对。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 取值路径，指向一个 JSON 对象
//
// 返回:
//   - map[string]gjson.Result: 对象的键值映射；路径不存在时返回空 map
func JSON_GetMap(jsonText string, path string) map[string]gjson.Result {
	return gjson.Get(jsonText, path).Map()
}

// JSON_Exists 判断 JSON 字符串中指定路径是否存在。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 检查路径
//
// 返回:
//   - bool: 路径存在返回 true，否则返回 false
func JSON_Exists(jsonText string, path string) bool {
	return gjson.Get(jsonText, path).Exists()
}

// JSON_ArrayLen 从 JSON 字符串中获取指定路径数组的长度。
//
// 参数:
//   - jsonText: JSON 格式的字符串
//   - path: 指向 JSON 数组的路径
//
// 返回:
//   - int: 数组长度；路径不存在或非数组时返回 0
func JSON_ArrayLen(jsonText string, path string) int {
	result := gjson.Get(jsonText, path)
	if result.IsArray() {
		return len(result.Array())
	}
	return 0
}

// JSON_Set 在 JSON 字符串中设置指定路径的值。
// 如果路径不存在会自动创建；支持设置字符串、数值、布尔值等。
//
// 参数:
//   - jsonText: 原始 JSON 字符串
//   - path: 设置路径，如 "user.name"、"list.-1"（追加到数组末尾）
//   - value: 要设置的值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func JSON_Set(jsonText string, path string, value interface{}) (string, error) {
	return sjson.Set(jsonText, path, value)
}

// JSON_SetString 在 JSON 字符串中设置指定路径的字符串值。
//
// 参数:
//   - jsonText: 原始 JSON 字符串
//   - path: 设置路径
//   - value: 要设置的字符串值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func JSON_SetString(jsonText string, path string, value string) (string, error) {
	return sjson.Set(jsonText, path, value)
}

// JSON_SetInt 在 JSON 字符串中设置指定路径的整数值。
//
// 参数:
//   - jsonText: 原始 JSON 字符串
//   - path: 设置路径
//   - value: 要设置的整数值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func JSON_SetInt(jsonText string, path string, value int64) (string, error) {
	return sjson.Set(jsonText, path, value)
}

// JSON_Delete 从 JSON 字符串中删除指定路径的值。
//
// 参数:
//   - jsonText: 原始 JSON 字符串
//   - path: 要删除的路径
//
// 返回:
//   - string: 删除后的 JSON 字符串
//   - error: 删除失败时返回错误
func JSON_Delete(jsonText string, path string) (string, error) {
	return sjson.Delete(jsonText, path)
}

// JSON_ToString 将 Go 值序列化为 JSON 字符串。
//
// 参数:
//   - v: 待序列化的值，支持结构体、map、切片等
//
// 返回:
//   - string: JSON 字符串
//   - error: 序列化失败时返回错误
func JSON_ToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// JSON_ToIndent 将 Go 值序列化为带缩进的 JSON 字符串。
//
// 参数:
//   - v: 待序列化的值
//   - indent: 缩进字符串，通常为 "  " 或 "\t"
//
// 返回:
//   - string: 格式化后的 JSON 字符串
//   - error: 序列化失败时返回错误
func JSON_ToIndent(v interface{}, indent string) (string, error) {
	data, err := json.MarshalIndent(v, "", indent)
	if err != nil {
		return "", err
	}
	return string(data), nil
}