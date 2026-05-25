// 类型转换工具
// 基于 spf13/cast 库，提供任意类型到基本类型的转换。
// 支持字符串、整数、浮点数、布尔值、时间、切片等类型的互转。
package utils

import (
	"strconv"

	"github.com/spf13/cast"
)

// Cast_ToString 将任意类型的值转换为字符串。
// 基于 cast.ToStringE 实现，支持基本类型及常见复合类型的转换。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - string: 转换后的字符串
func Cast_ToString(v interface{}) string {
	return cast.ToString(v)
}

// Cast_ToInt 将任意类型的值转换为 int。
// 字符串 "123" → 123，浮点数 3.14 → 3，布尔值 true → 1。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - int: 转换后的整数
func Cast_ToInt(v interface{}) int {
	return cast.ToInt(v)
}

// Cast_ToInt64 将任意类型的值转换为 int64。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - int64: 转换后的 int64 值
func Cast_ToInt64(v interface{}) int64 {
	return cast.ToInt64(v)
}

// Cast_ToFloat64 将任意类型的值转换为 float64。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - float64: 转换后的浮点数
func Cast_ToFloat64(v interface{}) float64 {
	return cast.ToFloat64(v)
}

// Cast_ToBool 将任意类型的值转换为 bool。
// 字符串 "true"/"1" → true，"false"/"0" → false。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - bool: 转换后的布尔值
func Cast_ToBool(v interface{}) bool {
	return cast.ToBool(v)
}

// Cast_ToStringSlice 将任意类型的值转换为 []string。
// 支持逗号分隔的字符串 "a,b,c" → ["a","b","c"]，
// 也支持 []interface{} → []string。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - []string: 字符串切片
func Cast_ToStringSlice(v interface{}) []string {
	return cast.ToStringSlice(v)
}

// Cast_ToIntSlice 将任意类型的值转换为 []int。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - []int: 整数切片
func Cast_ToIntSlice(v interface{}) []int {
	return cast.ToIntSlice(v)
}

// Cast_ToTime 将任意类型的值转换为 time.Time。
// 支持常见的时间格式字符串，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - interface{}: 转换后的时间（返回 interface{} 以保证兼容）
func Cast_ToTime(v interface{}) interface{} {
	return cast.ToTime(v)
}

// Cast_ToDuration 将任意类型的值转换为 time.Duration。
// 字符串 "5s" → 5秒，"1h30m" → 1小时30分钟。
//
// 参数:
//   - v: 待转换的值
//
// 返回:
//   - interface{}: 转换后的时间间隔（返回 interface{} 以保证兼容）
func Cast_ToDuration(v interface{}) interface{} {
	return cast.ToDuration(v)
}

// Cast_SafeToString 将任意类型的值转换为字符串，转换失败时返回默认值。
//
// 参数:
//   - v: 待转换的值
//   - defaultVal: 转换失败时返回的默认字符串
//
// 返回:
//   - string: 转换后的字符串或默认值
func Cast_SafeToString(v interface{}, defaultVal string) string {
	result, err := cast.ToStringE(v)
	if err != nil {
		return defaultVal
	}
	return result
}

// Cast_SafeToInt 将任意类型的值转换为 int，转换失败时返回默认值。
//
// 参数:
//   - v: 待转换的值
//   - defaultVal: 转换失败时返回的默认整数
//
// 返回:
//   - int: 转换后的整数或默认值
func Cast_SafeToInt(v interface{}, defaultVal int) int {
	result, err := cast.ToIntE(v)
	if err != nil {
		return defaultVal
	}
	return result
}

// Cast_SafeToFloat64 将任意类型的值转换为 float64，转换失败时返回默认值。
//
// 参数:
//   - v: 待转换的值
//   - defaultVal: 转换失败时返回的默认浮点数
//
// 返回:
//   - float64: 转换后的浮点数或默认值
func Cast_SafeToFloat64(v interface{}, defaultVal float64) float64 {
	result, err := cast.ToFloat64E(v)
	if err != nil {
		return defaultVal
	}
	return result
}

// Cast_SafeToBool 将任意类型的值转换为 bool，转换失败时返回默认值。
//
// 参数:
//   - v: 待转换的值
//   - defaultVal: 转换失败时返回的默认布尔值
//
// 返回:
//   - bool: 转换后的布尔值或默认值
func Cast_SafeToBool(v interface{}, defaultVal bool) bool {
	result, err := cast.ToBoolE(v)
	if err != nil {
		return defaultVal
	}
	return result
}

// Cast_ParseInt 将字符串按指定进制转换为 int64。
// 如 "ff" 按16进制 → 255，"1010" 按2进制 → 10。
//
// 参数:
//   - s: 待转换的字符串
//   - base: 进制（2/8/10/16）
//
// 返回:
//   - int64: 转换后的整数值
//   - error: 转换失败时返回错误
func Cast_ParseInt(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}