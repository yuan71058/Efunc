package utils

import (
	"regexp"
	"strconv"
)

// Z正则_校验密码 校验密码强度，要求 5-17 个非空白字符。
// 如果不匹配，通过 msg 指针返回错误提示信息。
//
// 参数:
//   - s: 待校验的密码字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示密码格式合法
func Z正则_校验密码(s string, msg *string) bool {
	匹配结果, _ := regexp.MatchString(`\S{5,17}$`, s)

	if !匹配结果 {
		*msg = "以字母开头，长度在6-18之间，只能包含字符、数字和下划线"
	}
	return 匹配结果
}

// Z正则_校验代理用户名 校验代理用户名，只允许英文字母、数字和中文字符。
//
// 参数:
//   - s: 待校验的用户名字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示用户名格式合法
func Z正则_校验代理用户名(s string, msg *string) bool {
	匹配结果, _ := regexp.MatchString(`^[a-zA-Z0-9\p{Han}]+$`, s)
	if !匹配结果 {
		*msg = "只能包含英文字母、数字和ANSI编码支持的中文字符"
	}
	return 匹配结果
}

// Z正则_校验用户名 校验用户名，要求 5-17 个单词字符（字母、数字、下划线）。
//
// 参数:
//   - s: 待校验的用户名字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示用户名格式合法
func Z正则_校验用户名(s string, msg *string) bool {
	匹配结果, _ := regexp.MatchString(`\w{5,17}$`, s)
	if 匹配结果 {
		return 匹配结果
	}
	*msg = "长度在6-18之间，只能包含字符、数字和下划线"
	return 匹配结果
}

// Z正则_校验email 校验电子邮件地址格式。
// 支持常见的邮箱格式，如 user@example.com、user.name+tag@sub.domain.com。
//
// 参数:
//   - s: 待校验的邮箱字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示邮箱格式合法
func Z正则_校验email(s string, msg *string) bool {
	匹配结果, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, s)
	if 匹配结果 {
		return 匹配结果
	}
	*msg = "非正确email格式"
	return 匹配结果
}

// Z正则_校验纯数字 校验字符串是否为纯数字（支持负号和小数点）。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示是纯数字
func Z正则_校验纯数字(s string, msg *string) bool {
	匹配结果, _ := regexp.MatchString(`^-?\d*\.?\d+$`, s)
	if 匹配结果 {
		return 匹配结果
	}
	*msg = "非完全是数字"
	return 匹配结果
}

// Z正则_校验纯数字指定位数 校验字符串是否为指定位数的纯数字。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//   - 位数: 要求的数字位数
//
// 返回:
//   - bool: true 表示符合指定位数的纯数字
func Z正则_校验纯数字指定位数(s string, msg *string, 位数 int) bool {
	匹配结果, _ := regexp.MatchString(`^\d{`+strconv.Itoa(位数)+`}$`, s)

	if 匹配结果 {
		return 匹配结果
	}
	*msg = "长度不为" + strconv.Itoa(位数)
	return 匹配结果
}

// Z正则_是否英数 校验字符串是否仅包含英文字母和数字。
//
// 参数:
//   - s: 待校验的字符串
//   - msg: 校验失败时的错误信息输出指针
//
// 返回:
//   - bool: true 表示仅包含英文字母和数字
func Z正则_是否英数(s string, msg *string) bool {
	reg := regexp.MustCompile("^[A-Za-z0-9]+$")
	if reg.MatchString(s) {
		return true
	}
	*msg = "只能输入数字字母"
	return false
}

// Z正则_取Url连接地址 从文本中提取所有 HTTP/HTTPS URL 链接。
//
// 参数:
//   - str: 包含 URL 的源文本
//
// 返回:
//   - []string: 提取到的 URL 数组
func Z正则_取Url连接地址(str string) []string {
	urlRegex := regexp.MustCompile(`(https?://[^\s]+)`)
	urls := urlRegex.FindAllString(str, -1)
	return urls
}

// Z正则_取全部匹配子文本 使用正则表达式提取所有匹配的子文本。
//
// 参数:
//   - str: 源文本
//   - 正则表达式: 正则表达式模式
//
// 返回:
//   - []string: 所有匹配结果的数组
func Z正则_取全部匹配子文本(str, 正则表达式 string) []string {
	urlRegex := regexp.MustCompile(正则表达式)
	urls := urlRegex.FindAllString(str, -1)
	return urls
}

// Z正则_取ip端口 从文本中提取第一个 IP:端口 格式的地址。
//
// 参数:
//   - str: 包含 IP:端口的源文本
//
// 返回:
//   - string: 第一个匹配的 IP:端口 字符串；未找到返回空串
func Z正则_取ip端口(str string) string {
	urlRegex := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+:\d+)`)
	urls := urlRegex.FindAllString(str, -1)
	if len(urls) == 0 {
		return ""
	}
	return urls[0]
}

// Z正则_取ip端口多个 从文本中提取所有 IP:端口 格式的地址。
//
// 参数:
//   - str: 包含 IP:端口的源文本
//
// 返回:
//   - []string: 所有匹配的 IP:端口 字符串数组；未找到返回空数组
func Z正则_取ip端口多个(str string) []string {
	urlRegex := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+:\d+)`)
	urls := urlRegex.FindAllString(str, -1)
	if len(urls) == 0 {
		return []string{}
	}
	return urls
}
