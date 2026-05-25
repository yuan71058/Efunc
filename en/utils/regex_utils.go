package utils

import (
	"regexp"
	"strconv"
)

// Regex_CheckPassword 校验密码强度，要求 5-17 个非空白字符。
// 如果不匹配，通过 msg 指针返回错误提示信息。
//
// 参数:
//   - s: 待校验的密码字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示密码格式合法
func Regex_CheckPassword(s string, msg *string) bool {
	match, _ := regexp.MatchString(`\S{5,17}$`, s)
	if !match {
		*msg = "以字母开头，长度在6-18之间，只能包含字符、数字和下划线"
	}
	return match
}

// Regex_CheckProxyUsername 校验代理用户名，只允许英文字母、数字和中文字符。
//
// 参数:
//   - s: 待校验的用户名字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示用户名格式合法
func Regex_CheckProxyUsername(s string, msg *string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9\p{Han}]+$`, s)
	if !match {
		*msg = "只能包含英文字母、数字和ANSI编码支持的中文字符"
	}
	return match
}

// Regex_CheckUsername 校验用户名，要求 5-17 个单词字符（字母、数字、下划线）。
//
// 参数:
//   - s: 待校验的用户名字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示用户名格式合法
func Regex_CheckUsername(s string, msg *string) bool {
	match, _ := regexp.MatchString(`\w{5,17}$`, s)
	if match {
		return match
	}
	*msg = "长度在6-18之间，只能包含字符、数字和下划线"
	return match
}

// Regex_CheckEmail 校验电子邮件地址格式。
// 支持常见的邮箱格式，如 user@example.com、user.name+tag@sub.domain.com。
//
// 参数:
//   - s: 待校验的邮箱字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示邮箱格式合法
func Regex_CheckEmail(s string, msg *string) bool {
	match, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, s)
	if match {
		return match
	}
	*msg = "非正确email格式"
	return match
}

// Regex_CheckNumeric 校验字符串是否为纯数字（支持负号和小数点）。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示是纯数字
func Regex_CheckNumeric(s string, msg *string) bool {
	match, _ := regexp.MatchString(`^-?\d*\.?\d+$`, s)
	if match {
		return match
	}
	*msg = "非完全是数字"
	return match
}

// Regex_CheckDigitsN 校验字符串是否为指定位数的纯数字。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//   - n: 要求的数字位数
//
// 返回:
//   - bool: true 表示符合指定位数的纯数字
func Regex_CheckDigitsN(s string, msg *string, n int) bool {
	match, _ := regexp.MatchString(`^\d{`+strconv.Itoa(n)+`}$`, s)
	if match {
		return match
	}
	*msg = "长度不为" + strconv.Itoa(n)
	return match
}

// Regex_IsAlphaNumeric 校验字符串是否仅包含英文字母和数字。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示仅包含英文字母和数字
func Regex_IsAlphaNumeric(s string, msg *string) bool {
	reg := regexp.MustCompile("^[A-Za-z0-9]+$")
	if reg.MatchString(s) {
		return true
	}
	*msg = "只能输入数字字母"
	return false
}

// Regex_ExtractURLs 从文本中提取所有 HTTP/HTTPS URL 链接。
//
// 参数:
//   - str: 包含 URL 的源文本
//
// 返回:
//   - []string: 提取到的 URL 数组
func Regex_ExtractURLs(str string) []string {
	urlRegex := regexp.MustCompile(`(https?://[^\s]+)`)
	urls := urlRegex.FindAllString(str, -1)
	return urls
}

// Regex_FindAllMatches 使用正则表达式提取所有匹配的子文本。
//
// 参数:
//   - str: 源文本
//   - pattern: 正则表达式模式
//
// 返回:
//   - []string: 所有匹配结果的数组
func Regex_FindAllMatches(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(str, -1)
	return matches
}

// Regex_ExtractIPPort 从文本中提取第一个 IP:端口 格式的地址。
//
// 参数:
//   - str: 包含 IP:端口的源文本
//
// 返回:
//   - string: 第一个匹配的 IP:端口 字符串；未找到返回空串
func Regex_ExtractIPPort(str string) string {
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+:\d+)`)
	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return ""
	}
	return matches[0]
}

// Regex_ExtractAllIPPort 从文本中提取所有 IP:端口 格式的地址。
//
// 参数:
//   - str: 包含 IP:端口的源文本
//
// 返回:
//   - []string: 所有匹配的 IP:端口 字符串数组；未找到返回空数组
func Regex_ExtractAllIPPort(str string) []string {
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+:\d+)`)
	matches := re.FindAllString(str, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return matches
}