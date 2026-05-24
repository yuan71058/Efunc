package utils

import (
	"math/rand"
	"strings"
)

// F_选择 三元选择器（非泛型版），根据逻辑值返回两个参数中的一个。
// 推荐使用泛型版 S三元 替代，类型更安全。
//
// 参数:
//   - 逻辑: 条件值，true 返回真返回参数，false 返回假返回参数
//   - 真返回参数: 条件为真时返回的值
//   - 假返回参数: 条件为假时返回的值
//
// 返回:
//   - interface{}: 选中的值
func F_选择(逻辑 bool, 真返回参数, 假返回参数 interface{}) interface{} {
	if 逻辑 {
		return 真返回参数
	}
	return 假返回参数
}

// F_取随机数 生成指定范围内的随机整数，包含边界值 [min, max]。
//
// 参数:
//   - min: 最小值（包含）
//   - max: 最大值（包含）
//
// 返回:
//   - int: [min, max] 范围内的随机整数
func F_取随机数(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// F_取文本右边 取文本右侧指定字节数的子串。
// 注意：按字节截取，中文可能被截断。中文安全截取请用 W文本_取右边。
//
// 参数:
//   - text: 源文本
//   - n: 截取的字节数
//
// 返回:
//   - string: 右侧 n 个字节的文本；n 超过长度时返回原文本
func F_取文本右边(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[len(text)-n:]
}

// F_取文本左边 取文本左侧指定字节数的子串。
// 注意：按字节截取，中文可能被截断。中文安全截取请用 W文本_取左边。
//
// 参数:
//   - text: 源文本
//   - n: 截取的字节数
//
// 返回:
//   - string: 左侧 n 个字节的文本；n 超过长度时返回原文本
func F_取文本左边(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[:n]
}

// F_加入成员 向字符串数组追加一个成员，返回新数组。
// 等同于 append(数组, 成员)。
//
// 参数:
//   - 数组: 原字符串数组
//   - 成员: 要追加的字符串
//
// 返回:
//   - []string: 追加后的新数组
func F_加入成员(数组 []string, 成员 string) []string {
	return append(数组, 成员)
}

// F_删首尾空 去除文本首尾的空白字符（空格、制表符、换行等）。
// 等同于 strings.TrimSpace。
//
// 参数:
//   - text: 源文本
//
// 返回:
//   - string: 去除首尾空白后的文本
func F_删首尾空(text string) string {
	return strings.TrimSpace(text)
}

// F_取文本长度 获取文本的字节长度。
// 注意：中文在 UTF-8 编码下占 3 个字节。如需字符数请用 W文本_取长度。
//
// 参数:
//   - text: 源文本
//
// 返回:
//   - int: 字节长度
func F_取文本长度(text string) int {
	return len(text)
}

// F_分割文本 按指定分隔符将文本分割为字符串数组。
// 等同于 strings.Split。
//
// 参数:
//   - 原文本: 待分割的文本
//   - 分割符: 用作分割的分隔符
//
// 返回:
//   - []string: 分割后的字符串数组
func F_分割文本(原文本 string, 分割符 string) []string {
	return strings.Split(原文本, 分割符)
}
