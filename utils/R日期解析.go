package utils

import (
	"time"

	"github.com/araddon/dateparse"
)

// R日期_智能解析 自动识别日期时间字符串的格式并解析为 time.Time。
// 支持绝大多数常见的日期时间格式，无需预先指定格式字符串。
// 基于 araddon/dateparse 库实现，支持中文日期、RFC3339、ISO8601 等格式。
//
// 参数:
//   - 日期文本: 日期时间字符串，如 "2024-01-15"、"2024/01/15 10:30:00"、"Jan 15, 2024"
//
// 返回:
//   - time.Time: 解析后的时间对象
//   - error: 解析失败时返回错误信息
func R日期_智能解析(日期文本 string) (time.Time, error) {
	return dateparse.ParseAny(日期文本)
}

// R日期_智能解析本地 自动识别日期时间字符串格式，解析为本地时区的时间。
// 与 R日期_智能解析 类似，但返回本地时区而非 UTC 时区的时间。
//
// 参数:
//   - 日期文本: 日期时间字符串
//
// 返回:
//   - time.Time: 本地时区的时间对象
//   - error: 解析失败时返回错误信息
func R日期_智能解析本地(日期文本 string) (time.Time, error) {
	return dateparse.ParseLocal(日期文本)
}

// R日期_智能解析严格 以严格模式解析日期时间字符串。
// 严格模式下，对格式匹配更加严格，减少误解析的可能性。
//
// 参数:
//   - 日期文本: 日期时间字符串
//
// 返回:
//   - time.Time: 解析后的时间对象
//   - error: 解析失败时返回错误信息
func R日期_智能解析严格(日期文本 string) (time.Time, error) {
	return dateparse.ParseStrict(日期文本)
}

// R日期_取格式名称 获取日期时间字符串对应的 Go 时间格式名称。
// 返回可用于 time.Parse 的格式字符串，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - 日期文本: 日期时间字符串
//
// 返回:
//   - string: Go 时间格式字符串；无法识别时返回空串
func R日期_取格式名称(日期文本 string) string {
	format, err := dateparse.ParseFormat(日期文本)
	if err != nil {
		return ""
	}
	return format
}

// R日期_到文本 将 time.Time 格式化为指定布局的字符串。
// 使用 Go 标准的时间布局格式，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - 时间: 时间对象
//   - 格式: Go 时间布局字符串，如 "2006-01-02"
//
// 返回:
//   - string: 格式化后的时间字符串
func R日期_到文本(时间 time.Time, 格式 string) string {
	return 时间.Format(格式)
}

// R日期_取时间戳 将 time.Time 转换为 Unix 时间戳（秒）。
//
// 参数:
//   - 时间: 时间对象
//
// 返回:
//   - int64: Unix 时间戳（秒）
func R日期_取时间戳(时间 time.Time) int64 {
	return 时间.Unix()
}

// R日期_取毫秒时间戳 将 time.Time 转换为 Unix 毫秒时间戳。
//
// 参数:
//   - 时间: 时间对象
//
// 返回:
//   - int64: Unix 毫秒时间戳
func R日期_取毫秒时间戳(时间 time.Time) int64 {
	return 时间.UnixMilli()
}

// R日期_取日期部分 从 time.Time 中提取日期部分（年月日）。
//
// 参数:
//   - 时间: 时间对象
//
// 返回:
//   - string: 日期部分字符串，格式 "2006-01-02"
func R日期_取日期部分(时间 time.Time) string {
	return 时间.Format("2006-01-02")
}

// R日期_取时间部分 从 time.Time 中提取时间部分（时分秒）。
//
// 参数:
//   - 时间: 时间对象
//
// 返回:
//   - string: 时间部分字符串，格式 "15:04:05"
func R日期_取时间部分(时间 time.Time) string {
	return 时间.Format("15:04:05")
}
