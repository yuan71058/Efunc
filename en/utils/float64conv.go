package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// Float64_Abs 返回浮点数的绝对值，使用 decimal 库防止精度丢失。
// 负数乘以 -1 转为正数，正数和零直接返回。
//
// 参数:
//   - value: 输入的浮点数
//
// 返回:
//   - float64: 绝对值
func Float64_Abs(value float64) float64 {
	var result float64
	if value < 0 {
		decVal := decimal.NewFromFloat(value)
		decMul := decimal.NewFromInt(-1)
		result, _ = decVal.Mul(decMul).Float64()
	} else {
		result = value
	}
	return result
}

// Float64_MulInt64 将 float64 与 int64 相乘，使用 decimal 库防止精度丢失。
//
// 参数:
//   - val1: 浮点乘数
//   - val2: 整数乘数
//
// 返回:
//   - float64: 乘积结果
func Float64_MulInt64(val1 float64, val2 int64) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decMul := decimal.NewFromInt(val2)
	result, _ = decVal.Mul(decMul).Float64()
	return result
}

// Float64_Mul 将两个 float64 相乘，使用 decimal 库防止精度丢失。
//
// 参数:
//   - val1: 第一个浮点乘数
//   - val2: 第二个浮点乘数
//
// 返回:
//   - float64: 乘积结果
func Float64_Mul(val1 float64, val2 float64) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decMul := decimal.NewFromFloat(val2)
	result, _ = decVal.Mul(decMul).Float64()
	return result
}

// Float64_DivInt64 将 float64 除以 int64，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - val1: 被除数
//   - val2: 除数
//   - precision: 小数点后保留的位数
//
// 返回:
//   - float64: 除法结果（四舍五入到指定小数位）
func Float64_DivInt64(val1 float64, val2 int64, precision int32) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decDiv := decimal.NewFromInt(val2)
	result, _ = decVal.Div(decDiv).Round(precision).Float64()
	return result
}

// Float64_Div 将两个 float64 相除，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - val1: 被除数
//   - val2: 除数
//   - precision: 小数点后保留的位数
//
// 返回:
//   - float64: 除法结果（四舍五入到指定小数位）
func Float64_Div(val1 float64, val2 float64, precision int32) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decDiv := decimal.NewFromFloat(val2)
	result, _ = decVal.Div(decDiv).Round(precision).Float64()
	return result
}

// Float64_Neg 返回浮点数的负值，使用 decimal 库防止精度丢失。
// 正数取反为负数，零和负数原样返回。
//
// 参数:
//   - value: 输入的浮点数
//
// 返回:
//   - float64: 负值结果
func Float64_Neg(value float64) float64 {
	var result float64
	if value > 0 {
		decVal := decimal.NewFromFloat(value)
		decMul := decimal.NewFromInt(-1)
		result, _ = decVal.Mul(decMul).Float64()
	} else {
		result = value
	}
	return result
}

// Float64_ToString 将 float64 转换为指定小数位数的字符串。
// 使用 strconv.FormatFloat 进行格式化。
//
// 参数:
//   - value: 待转换的浮点数
//   - precision: 小数点后保留的位数
//
// 返回:
//   - string: 格式化后的浮点数字符串
//
// 示例:
//
//	Float64_ToString(3.14159, 2)  // "3.14"
func Float64_ToString(value float64, precision int) string {
	return strconv.FormatFloat(value, 'f', precision, 64)
}

// Float64_FromString 将字符串转换为 float64，并按指定小数位四舍五入。
// 使用 decimal 库确保精度，转换失败返回 0。
//
// 参数:
//   - value: 待转换的字符串
//   - precision: 小数点后保留的位数
//
// 返回:
//   - float64: 转换后的浮点数；解析失败返回 0
func Float64_FromString(value string, precision int) float64 {
	decVal, err := decimal.NewFromString(value)
	if err != nil {
		return 0
	}
	result, _ := decVal.Round(int32(precision)).Float64()
	return result
}

// Int64ToFloat64 将 int64 转换为 float64，使用 decimal 库防止精度丢失。
// 结果保留 2 位小数。
//
// 参数:
//   - value: 待转换的 int64 值
//
// 返回:
//   - float64: 转换后的浮点数（保留 2 位小数）
func Int64ToFloat64(value int64) float64 {
	var result float64
	decVal := decimal.NewFromInt(value)
	result, _ = decVal.Round(2).Float64()
	return result
}

// Float64_Sub 两个 float64 相减，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - val1: 被减数
//   - val2: 减数
//   - precision: 小数点后保留的位数
//
// 返回:
//   - float64: 减法结果（四舍五入到指定小数位）
func Float64_Sub(val1 float64, val2 float64, precision int32) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decSub := decimal.NewFromFloat(val2)
	result, _ = decVal.Sub(decSub).Round(precision).Float64()
	return result
}

// Float64_Add 两个 float64 相加，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - val1: 第一个加数
//   - val2: 第二个加数
//   - precision: 小数点后保留的位数
//
// 返回:
//   - float64: 加法结果（四舍五入到指定小数位）
func Float64_Add(val1 float64, val2 float64, precision int32) float64 {
	var result float64
	decVal := decimal.NewFromFloat(val1)
	decAdd := decimal.NewFromFloat(val2)
	result, _ = decVal.Add(decAdd).Round(precision).Float64()
	return result
}