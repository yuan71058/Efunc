// 智能日期解析工具
// 基于 araddon/dateparse 库，自动识别日期时间字符串的格式并解析。
// 支持绝大多数常见的日期时间格式，无需预先指定格式字符串。
// 支持中文日期、RFC3339、ISO8601 等格式。
package utils

import (
	"time"

	"github.com/araddon/dateparse"
)

// DateParse_Any 自动识别日期时间字符串的格式并解析为 time.Time。
// 支持绝大多数常见的日期时间格式，无需预先指定格式字符串。
//
// 参数:
//   - dateText: 日期时间字符串，如 "2024-01-15"、"2024/01/15 10:30:00"、"Jan 15, 2024"
//
// 返回:
//   - time.Time: 解析后的时间对象
//   - error: 解析失败时返回错误信息
func DateParse_Any(dateText string) (time.Time, error) {
	return dateparse.ParseAny(dateText)
}

// DateParse_Local 自动识别日期时间字符串格式，解析为本地时区的时间。
// 与 DateParse_Any 类似，但返回本地时区而非 UTC 时区的时间。
//
// 参数:
//   - dateText: 日期时间字符串
//
// 返回:
//   - time.Time: 本地时区的时间对象
//   - error: 解析失败时返回错误信息
func DateParse_Local(dateText string) (time.Time, error) {
	return dateparse.ParseLocal(dateText)
}

// DateParse_Strict 以严格模式解析日期时间字符串。
// 严格模式下，对格式匹配更加严格，减少误解析的可能性。
//
// 参数:
//   - dateText: 日期时间字符串
//
// 返回:
//   - time.Time: 解析后的时间对象
//   - error: 解析失败时返回错误信息
func DateParse_Strict(dateText string) (time.Time, error) {
	return dateparse.ParseStrict(dateText)
}

// DateParse_Format 获取日期时间字符串对应的 Go 时间格式名称。
// 返回可用于 time.Parse 的格式字符串，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - dateText: 日期时间字符串
//
// 返回:
//   - string: Go 时间格式字符串；无法识别时返回空串
func DateParse_Format(dateText string) string {
	format, err := dateparse.ParseFormat(dateText)
	if err != nil {
		return ""
	}
	return format
}

// DateParse_ToString 将 time.Time 格式化为指定布局的字符串。
// 使用 Go 标准的时间布局格式，如 "2006-01-02 15:04:05"。
//
// 参数:
//   - t: 时间对象
//   - layout: Go 时间布局字符串，如 "2006-01-02"
//
// 返回:
//   - string: 格式化后的时间字符串
func DateParse_ToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// DateParse_ToTimestamp 将 time.Time 转换为 Unix 时间戳（秒）。
//
// 参数:
//   - t: 时间对象
//
// 返回:
//   - int64: Unix 时间戳（秒）
func DateParse_ToTimestamp(t time.Time) int64 {
	return t.Unix()
}

// DateParse_ToMillis 将 time.Time 转换为 Unix 毫秒时间戳。
//
// 参数:
//   - t: 时间对象
//
// 返回:
//   - int64: Unix 毫秒时间戳
func DateParse_ToMillis(t time.Time) int64 {
	return t.UnixMilli()
}

// DateParse_GetDate 从 time.Time 中提取日期部分（年月日）。
//
// 参数:
//   - t: 时间对象
//
// 返回:
//   - string: 日期部分字符串，格式 "2006-01-02"
func DateParse_GetDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// DateParse_GetTime 从 time.Time 中提取时间部分（时分秒）。
//
// 参数:
//   - t: 时间对象
//
// 返回:
//   - string: 时间部分字符串，格式 "15:04:05"
func DateParse_GetTime(t time.Time) string {
	return t.Format("15:04:05")
}