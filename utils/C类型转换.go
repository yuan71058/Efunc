package utils

import (
	"strconv"

	"github.com/spf13/cast"
)

// C类型_到文本 将任意类型的值转换为字符串。
// 基于 cast.ToStringE 实现，支持基本类型及常见复合类型的转换。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - string: 转换后的字符串
func C类型_到文本(值 interface{}) string {
	return cast.ToString(值)
}

// C类型_到整数 将任意类型的值转换为 int。
// 字符串 "123" → 123，浮点数 3.14 → 3，布尔值 true → 1。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - int: 转换后的整数
func C类型_到整数(值 interface{}) int {
	return cast.ToInt(值)
}

// C类型_到整数64 将任意类型的值转换为 int64。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - int64: 转换后的 int64 值
func C类型_到整数64(值 interface{}) int64 {
	return cast.ToInt64(值)
}

// C类型_到浮点数 将任意类型的值转换为 float64。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - float64: 转换后的浮点数
func C类型_到浮点数(值 interface{}) float64 {
	return cast.ToFloat64(值)
}

// C类型_到逻辑型 将任意类型的值转换为 bool。
// 字符串 "true"/"1" → true，"false"/"0" → false。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - bool: 转换后的布尔值
func C类型_到逻辑型(值 interface{}) bool {
	return cast.ToBool(值)
}

// C类型_到文本切片 将任意类型的值转换为 []string。
// 支持逗号分隔的字符串 "a,b,c" → ["a","b","c"]，
// 也支持 []interface{} → []string。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - []string: 字符串切片
func C类型_到文本切片(值 interface{}) []string {
	return cast.ToStringSlice(值)
}

// C类型_到整数切片 将任意类型的值转换为 []int。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - []int: 整数切片
func C类型_到整数切片(值 interface{}) []int {
	return cast.ToIntSlice(值)
}

// C类型_到时间 将任意类型的值转换为 time.Time。
// 支持常见的时间格式字符串，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - time.Time: 转换后的时间
func C类型_到时间(值 interface{}) interface{} {
	return cast.ToTime(值)
}

// C类型_到Duration 将任意类型的值转换为 time.Duration。
// 字符串 "5s" → 5秒，"1h30m" → 1小时30分钟。
//
// 参数:
//   - 值: 待转换的值
//
// 返回:
//   - time.Duration: 转换后的时间间隔
func C类型_到Duration(值 interface{}) interface{} {
	return cast.ToDuration(值)
}

// C类型_安全到文本 将任意类型的值转换为字符串，转换失败时返回默认值。
//
// 参数:
//   - 值: 待转换的值
//   - 默认值: 转换失败时返回的默认字符串
//
// 返回:
//   - string: 转换后的字符串或默认值
func C类型_安全到文本(值 interface{}, 默认值 string) string {
	result, err := cast.ToStringE(值)
	if err != nil {
		return 默认值
	}
	return result
}

// C类型_安全到整数 将任意类型的值转换为 int，转换失败时返回默认值。
//
// 参数:
//   - 值: 待转换的值
//   - 默认值: 转换失败时返回的默认整数
//
// 返回:
//   - int: 转换后的整数或默认值
func C类型_安全到整数(值 interface{}, 默认值 int) int {
	result, err := cast.ToIntE(值)
	if err != nil {
		return 默认值
	}
	return result
}

// C类型_安全到浮点数 将任意类型的值转换为 float64，转换失败时返回默认值。
//
// 参数:
//   - 值: 待转换的值
//   - 默认值: 转换失败时返回的默认浮点数
//
// 返回:
//   - float64: 转换后的浮点数或默认值
func C类型_安全到浮点数(值 interface{}, 默认值 float64) float64 {
	result, err := cast.ToFloat64E(值)
	if err != nil {
		return 默认值
	}
	return result
}

// C类型_安全到逻辑型 将任意类型的值转换为 bool，转换失败时返回默认值。
//
// 参数:
//   - 值: 待转换的值
//   - 默认值: 转换失败时返回的默认布尔值
//
// 返回:
//   - bool: 转换后的布尔值或默认值
func C类型_安全到逻辑型(值 interface{}, 默认值 bool) bool {
	result, err := cast.ToBoolE(值)
	if err != nil {
		return 默认值
	}
	return result
}

// C类型_进制转换 将字符串按指定进制转换为 int64。
// 如 "ff" 按16进制 → 255，"1010" 按2进制 → 10。
//
// 参数:
//   - 文本: 待转换的字符串
//   - 进制: 进制（2/8/10/16）
//
// 返回:
//   - int64: 转换后的整数值
//   - error: 转换失败时返回错误
func C类型_进制转换(文本 string, 进制 int) (int64, error) {
	return strconv.ParseInt(文本, 进制, 64)
}
