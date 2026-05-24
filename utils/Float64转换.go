package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// Float64取绝对值 返回浮点数的绝对值，使用 decimal 库防止精度丢失。
// 负数乘以 -1 转为正数，正数和零直接返回。
//
// 参数:
//   - 值: 输入的浮点数
//
// 返回:
//   - float64: 绝对值
func Float64取绝对值(值 float64) float64 {
	var 最终 float64
	if 值 < 0 {
		局_精确 := decimal.NewFromFloat(值)
		局_精确乘数 := decimal.NewFromInt(-1)
		最终, _ = 局_精确.Mul(局_精确乘数).Float64()
	} else {
		最终 = 值
	}
	return 最终
}

// Float64乘int64 将 float64 与 int64 相乘，使用 decimal 库防止精度丢失。
//
// 参数:
//   - 值1: 浮点乘数
//   - 值2: 整数乘数
//
// 返回:
//   - float64: 乘积结果
func Float64乘int64(值1 float64, 值2 int64) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确乘数 := decimal.NewFromInt(值2)
	最终, _ = 局_精确.Mul(局_精确乘数).Float64()

	return 最终
}

// Float64乘Float64 将两个 float64 相乘，使用 decimal 库防止精度丢失。
//
// 参数:
//   - 值1: 第一个浮点乘数
//   - 值2: 第二个浮点乘数
//
// 返回:
//   - float64: 乘积结果
func Float64乘Float64(值1 float64, 值2 float64) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确乘数 := decimal.NewFromFloat(值2)
	最终, _ = 局_精确.Mul(局_精确乘数).Float64()

	return 最终
}

// Float64除int64 将 float64 除以 int64，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - 值1: 被除数
//   - 值2: 除数
//   - 保留长度: 小数点后保留的位数
//
// 返回:
//   - float64: 除法结果（四舍五入到指定小数位）
func Float64除int64(值1 float64, 值2 int64, 保留长度 int32) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确除数 := decimal.NewFromInt(值2)
	最终, _ = 局_精确.Div(局_精确除数).Round(保留长度).Float64()

	return 最终
}

// Float64除float64 将两个 float64 相除，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - 值1: 被除数
//   - 值2: 除数
//   - 保留长度: 小数点后保留的位数
//
// 返回:
//   - float64: 除法结果（四舍五入到指定小数位）
func Float64除float64(值1 float64, 值2 float64, 保留长度 int32) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确除数 := decimal.NewFromFloat(值2)
	最终, _ = 局_精确.Div(局_精确除数).Round(保留长度).Float64()

	return 最终
}

// Float64取负值 返回浮点数的负值，使用 decimal 库防止精度丢失。
// 正数取反为负数，零和负数原样返回。
//
// 参数:
//   - 值: 输入的浮点数
//
// 返回:
//   - float64: 负值结果
func Float64取负值(值 float64) float64 {
	var 最终 float64
	if 值 > 0 {
		局_精确 := decimal.NewFromFloat(值)
		局_精确乘数 := decimal.NewFromInt(-1)
		最终, _ = 局_精确.Mul(局_精确乘数).Float64()
	} else {
		最终 = 值
	}
	return 最终
}

// Float64到文本 将 float64 转换为指定小数位数的字符串。
// 使用 strconv.FormatFloat 进行格式化。
//
// 参数:
//   - 值: 待转换的浮点数
//   - 保留小数点多少位: 小数点后保留的位数
//
// 返回:
//   - string: 格式化后的浮点数字符串
//
// 示例:
//
//	Float64到文本(3.14159, 2)  // "3.14"
func Float64到文本(值 float64, 保留小数点多少位 int) string {

	return strconv.FormatFloat(值, 'f', 保留小数点多少位, 64)
}

// Float64从文本 将字符串转换为 float64，并按指定小数位四舍五入。
// 使用 decimal 库确保精度，转换失败返回 0。
//
// 参数:
//   - 值: 待转换的字符串
//   - 保留小数点多少位: 小数点后保留的位数
//
// 返回:
//   - float64: 转换后的浮点数；解析失败返回 0
func Float64从文本(值 string, 保留小数点多少位 int) float64 {
	局_精确, err := decimal.NewFromString(值)
	if err != nil {
		return 0
	}
	局_精确2, _ := 局_精确.Round(int32(保留小数点多少位)).Float64()

	return 局_精确2
}

// Int64到Float64 将 int64 转换为 float64，使用 decimal 库防止精度丢失。
// 结果保留 2 位小数。
//
// 参数:
//   - 值: 待转换的 int64 值
//
// 返回:
//   - float64: 转换后的浮点数（保留 2 位小数）
func Int64到Float64(值 int64) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromInt(值)
	最终, _ = 局_精确.Round(2).Float64()
	return 最终
}

// Float64减float64 两个 float64 相减，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - 值1: 被减数
//   - 值2: 减数
//   - 保留长度: 小数点后保留的位数
//
// 返回:
//   - float64: 减法结果（四舍五入到指定小数位）
func Float64减float64(值1 float64, 值2 float64, 保留长度 int32) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确除数 := decimal.NewFromFloat(值2)
	最终, _ = 局_精确.Sub(局_精确除数).Round(保留长度).Float64()
	return 最终
}

// Float64加float64 两个 float64 相加，使用 decimal 库防止精度丢失。
// 结果按指定保留长度进行四舍五入。
//
// 参数:
//   - 值1: 第一个加数
//   - 值2: 第二个加数
//   - 保留长度: 小数点后保留的位数
//
// 返回:
//   - float64: 加法结果（四舍五入到指定小数位）
func Float64加float64(值1 float64, 值2 float64, 保留长度 int32) float64 {
	var 最终 float64
	局_精确 := decimal.NewFromFloat(值1)
	局_精确除数 := decimal.NewFromFloat(值2)
	最终, _ = 局_精确.Add(局_精确除数).Round(保留长度).Float64()
	return 最终
}
