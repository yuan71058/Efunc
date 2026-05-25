package utils

import (
	"math/rand"
	"strings"
)

// Select 三元选择器（非泛型版），根据逻辑值返回两个参数中的一个。
// 推荐使用泛型版 Ternary 替代，类型更安全。
//
// 参数:
//   - condition: 条件值，true 返回 trueVal，false 返回 falseVal
//   - trueVal: 条件为真时返回的值
//   - falseVal: 条件为假时返回的值
//
// 返回:
//   - interface{}: 选中的值
func Select(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// RandomNum 生成指定范围内的随机整数，包含边界值 [min, max]。
//
// 参数:
//   - min: 最小值（包含）
//   - max: 最大值（包含）
//
// 返回:
//   - int: [min, max] 范围内的随机整数
func RandomNum(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// TextRight 取文本右侧指定字节数的子串。
// 注意：按字节截取，中文可能被截断。中文安全截取请用 Text_Right。
//
// 参数:
//   - text: 源文本
//   - n: 截取的字节数
//
// 返回:
//   - string: 右侧 n 个字节的文本；n 超过长度时返回原文本
func TextRight(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[len(text)-n:]
}

// TextLeft 取文本左侧指定字节数的子串。
// 注意：按字节截取，中文可能被截断。中文安全截取请用 Text_Left。
//
// 参数:
//   - text: 源文本
//   - n: 截取的字节数
//
// 返回:
//   - string: 左侧 n 个字节的文本；n 超过长度时返回原文本
func TextLeft(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[:n]
}

// AppendMember 向字符串数组追加一个成员，返回新数组。
// 等同于 append(arr, member)。
//
// 参数:
//   - arr: 原字符串数组
//   - member: 要追加的字符串
//
// 返回:
//   - []string: 追加后的新数组
func AppendMember(arr []string, member string) []string {
	return append(arr, member)
}

// Trim 去除文本首尾的空白字符（空格、制表符、换行等）。
// 等同于 strings.TrimSpace。
//
// 参数:
//   - text: 源文本
//
// 返回:
//   - string: 去除首尾空白后的文本
func Trim(text string) string {
	return strings.TrimSpace(text)
}

// TextLen 获取文本的字节长度。
// 注意：中文在 UTF-8 编码下占 3 个字节。如需字符数请用 Text_CharCount。
//
// 参数:
//   - text: 源文本
//
// 返回:
//   - int: 字节长度
func TextLen(text string) int {
	return len(text)
}

// SplitText 按指定分隔符将文本分割为字符串数组。
// 等同于 strings.Split。
//
// 参数:
//   - text: 待分割的文本
//   - separator: 用作分割的分隔符
//
// 返回:
//   - []string: 分割后的字符串数组
func SplitText(text string, separator string) []string {
	return strings.Split(text, separator)
}