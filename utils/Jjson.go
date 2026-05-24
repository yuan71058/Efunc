package utils

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Jjson_取值 从 JSON 字符串中根据路径获取值。
// 路径语法参考 gjson：点号分隔，如 "user.name"、"list.0.id"。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径，如 "key"、"a.b.c"、"arr.0"
//
// 返回:
//   - string: 路径对应的值（字符串形式）；路径不存在时返回空串
func Jjson_取值(json文本 string, 路径 string) string {
	return gjson.Get(json文本, 路径).String()
}

// Jjson_取整数 从 JSON 字符串中根据路径获取整数值。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径
//
// 返回:
//   - int64: 路径对应的整数值；路径不存在时返回 0
func Jjson_取整数(json文本 string, 路径 string) int64 {
	return gjson.Get(json文本, 路径).Int()
}

// Jjson_取浮点数 从 JSON 字符串中根据路径获取浮点数值。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径
//
// 返回:
//   - float64: 路径对应的浮点数值；路径不存在时返回 0
func Jjson_取浮点数(json文本 string, 路径 string) float64 {
	return gjson.Get(json文本, 路径).Float()
}

// Jjson_取逻辑型 从 JSON 字符串中根据路径获取布尔值。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径
//
// 返回:
//   - bool: 路径对应的布尔值；路径不存在时返回 false
func Jjson_取逻辑型(json文本 string, 路径 string) bool {
	return gjson.Get(json文本, 路径).Bool()
}

// Jjson_取数组 从 JSON 字符串中根据路径获取数组。
// 返回 gjson.Result 数组，可通过下标访问每个元素。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径，指向一个 JSON 数组
//
// 返回:
//   - []gjson.Result: 数组元素列表；路径不存在时返回空切片
func Jjson_取数组(json文本 string, 路径 string) []gjson.Result {
	return gjson.Get(json文本, 路径).Array()
}

// Jjson_取对象 从 JSON 字符串中根据路径获取对象的键值对。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 取值路径，指向一个 JSON 对象
//
// 返回:
//   - map[string]gjson.Result: 对象的键值映射；路径不存在时返回空 map
func Jjson_取对象(json文本 string, 路径 string) map[string]gjson.Result {
	return gjson.Get(json文本, 路径).Map()
}

// Jjson_是否存在 判断 JSON 字符串中指定路径是否存在。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 检查路径
//
// 返回:
//   - bool: 路径存在返回 true，否则返回 false
func Jjson_是否存在(json文本 string, 路径 string) bool {
	return gjson.Get(json文本, 路径).Exists()
}

// Jjson_取数组长度 从 JSON 字符串中获取指定路径数组的长度。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 路径: 指向 JSON 数组的路径
//
// 返回:
//   - int: 数组长度；路径不存在或非数组时返回 0
func Jjson_取数组长度(json文本 string, 路径 string) int {
	result := gjson.Get(json文本, 路径)
	if result.IsArray() {
		return len(result.Array())
	}
	return 0
}

// Jjson_设置值 在 JSON 字符串中设置指定路径的值。
// 如果路径不存在会自动创建；支持设置字符串、数值、布尔值等。
//
// 参数:
//   - json文本: 原始 JSON 字符串
//   - 路径: 设置路径，如 "user.name"、"list.-1"（追加到数组末尾）
//   - 值: 要设置的值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func Jjson_设置值(json文本 string, 路径 string, 值 interface{}) (string, error) {
	return sjson.Set(json文本, 路径, 值)
}

// Jjson_设置文本值 在 JSON 字符串中设置指定路径的字符串值。
//
// 参数:
//   - json文本: 原始 JSON 字符串
//   - 路径: 设置路径
//   - 值: 要设置的字符串值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func Jjson_设置文本值(json文本 string, 路径 string, 值 string) (string, error) {
	return sjson.Set(json文本, 路径, 值)
}

// Jjson_设置整数值 在 JSON 字符串中设置指定路径的整数值。
//
// 参数:
//   - json文本: 原始 JSON 字符串
//   - 路径: 设置路径
//   - 值: 要设置的整数值
//
// 返回:
//   - string: 修改后的 JSON 字符串
//   - error: 设置失败时返回错误
func Jjson_设置整数值(json文本 string, 路径 string, 值 int64) (string, error) {
	return sjson.Set(json文本, 路径, 值)
}

// Jjson_删除值 从 JSON 字符串中删除指定路径的值。
//
// 参数:
//   - json文本: 原始 JSON 字符串
//   - 路径: 要删除的路径
//
// 返回:
//   - string: 删除后的 JSON 字符串
//   - error: 删除失败时返回错误
func Jjson_删除值(json文本 string, 路径 string) (string, error) {
	return sjson.Delete(json文本, 路径)
}

// Jjson_到文本 将 Go 值序列化为 JSON 字符串。
//
// 参数:
//   - 值: 待序列化的值，支持结构体、map、切片等
//
// 返回:
//   - string: JSON 字符串
//   - error: 序列化失败时返回错误
func Jjson_到文本(值 interface{}) (string, error) {
	数据, 错误 := json.Marshal(值)
	if 错误 != nil {
		return "", 错误
	}
	return string(数据), nil
}

// Jjson_到格式化文本 将 Go 值序列化为带缩进的 JSON 字符串。
//
// 参数:
//   - 值: 待序列化的值
//   - 缩进: 缩进字符串，通常为 "  " 或 "\t"
//
// 返回:
//   - string: 格式化后的 JSON 字符串
//   - error: 序列化失败时返回错误
func Jjson_到格式化文本(值 interface{}, 缩进 string) (string, error) {
	数据, 错误 := json.MarshalIndent(值, "", 缩进)
	if 错误 != nil {
		return "", 错误
	}
	return string(数据), nil
}

// Jjson_解析 将 JSON 字符串反序列化到目标变量。
//
// 参数:
//   - json文本: JSON 格式的字符串
//   - 目标: 指向目标变量的指针，如 &obj
//
// 返回:
//   - error: 反序列化失败时返回错误
func Jjson_解析(json文本 string, 目标 interface{}) error {
	return json.Unmarshal([]byte(json文本), 目标)
}

// Jjson_取所有路径 获取 JSON 字符串中所有值的路径。
// 返回所有叶节点的路径，便于遍历整个 JSON 结构。
//
// 参数:
//   - json文本: JSON 格式的字符串
//
// 返回:
//   - []string: 所有路径的切片
func Jjson_取所有路径(json文本 string) []string {
	var 路径列表 []string
	gjson.Parse(json文本).ForEach(func(key, value gjson.Result) bool {
		路径列表 = append(路径列表, key.String())
		if value.IsObject() || value.IsArray() {
			子路径 := Jjson_取所有路径(value.String())
			for _, p := range 子路径 {
				路径列表 = append(路径列表, key.String()+"."+p)
			}
		}
		return true
	})
	return 路径列表
}
