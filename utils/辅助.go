package utils

import (
	"math/rand"
	"strings"
)

func 选择(逻辑 bool, 真返回参数, 假返回参数 interface{}) interface{} {
	if 逻辑 {
		return 真返回参数
	}
	return 假返回参数
}
func 取随机数(min, max int) int {
	return rand.Intn(max-min+1) + min
}
func 取文本右边(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[len(text)-n:]
}

func 取文本左边(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[:n]
}
func 加入成员(数组 []string, 成员 string) []string {
	return append(数组, 成员)
}
func 删首尾空(text string) string {
	return strings.TrimSpace(text)
}
func 取文本长度(text string) int {
	return len(text)
}

func 分割文本(原文本 string, 分割符 string) []string {
	return strings.Split(原文本, 分割符)
}
