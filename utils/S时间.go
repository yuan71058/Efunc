package utils

import (
	"fmt"
	"strings"
	"time"
)

// base_format 默认的时间格式模板，对应 Go 的 "2006-01-02 15:04:05"
const base_format = "2006-01-02 15:04:05"

// S时间_文本到时间戳 将时间文本转换为 10 位 Unix 时间戳（秒）。
// 使用本地时区解析，格式必须为 "2006-01-02 15:04:05"。
//
// 参数:
//   - 时间文本: 格式为 "2006-01-02 15:04:05" 的时间字符串
//
// 返回:
//   - int: 10 位 Unix 时间戳；解析失败返回 0
func S时间_文本到时间戳(时间文本 string) int {

	formatTime, err := time.ParseInLocation("2006-01-02 15:04:05", 时间文本, time.Local)
	if err == nil {
		return int(formatTime.Unix())
	}
	return 0

}

// S时间_取现行时间戳13 获取当前时间的 13 位毫秒级时间戳。
//
// 返回:
//   - int64: 13 位毫秒级时间戳
func S时间_取现行时间戳13() int64 {
	return time.Now().UnixNano() / 1e6
}

// S时间_取现行时间戳 获取当前时间的 10 位秒级时间戳。
//
// 返回:
//   - int64: 10 位秒级时间戳
func S时间_取现行时间戳() int64 {
	return time.Now().Unix()
}

// S时间_取现行时间 获取当前系统日期和时间，格式为 "2006-01-02 15:04:05"。
//
// 返回:
//   - string: 格式化的当前时间字符串
func S时间_取现行时间() string {
	t := time.Now()
	str_time := t.Format(base_format)
	return str_time
}

// S时间_时间戳到时间 将 10 位秒级时间戳转换为格式化的时间字符串。
//
// 参数:
//   - 时间戳: 10 位秒级时间戳
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的时间字符串
func S时间_时间戳到时间(时间戳 int64) string {
	return time.Unix(时间戳, 0).Format(base_format)
}

// S时间_时间戳13到时间 将 13 位毫秒级时间戳转换为格式化的时间字符串。
//
// 参数:
//   - 时间戳: 13 位毫秒级时间戳
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的时间字符串
func S时间_时间戳13到时间(时间戳 int64) string {
	return time.Unix(时间戳/1000, 0).Format(base_format)
}

// S时间_时间到时间戳 将格式化的时间字符串转换为 10 位秒级时间戳。
// 使用 UTC 时区解析，与 S时间_文本到时间戳（本地时区）可能有时区差异。
//
// 参数:
//   - 时间: 格式为 "2006-01-02 15:04:05" 的时间字符串
//
// 返回:
//   - int64: 10 位秒级时间戳；解析失败返回 0
func S时间_时间到时间戳(时间 string) int64 {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, 时间)
	if err != nil {
		return 0
	}
	timestampInSeconds := t.Unix()
	return timestampInSeconds
}

// S时间_时间戳格式化 使用自定义格式字符串格式化时间戳。
// 支持的占位符：y/Y=年, m/M=月, d/D=日, h=12小时制, H=24小时制, i=分, s=秒, t=am/pm, T=AM/PM。
// 时间戳为 0 时使用当前时间。
//
// 参数:
//   - format: 自定义格式字符串（如 "Y-m-d H:i:s"）
//   - 时间戳: 10 位秒级时间戳，0 表示当前时间
//
// 返回:
//   - string: 格式化后的时间字符串
func S时间_时间戳格式化(format string, 时间戳 int64) string {
	var tm time.Time

	if 时间戳 == 0 {
		tm = time.Now()
	} else {
		tm = time.Unix(时间戳, 0)
	}

	patterns := []string{
		"y", "2006",
		"m", "01",
		"d", "02",

		"Y", "2006",
		"M", "01",
		"D", "02",

		"h", "3",
		"H", "15",

		"i", "04",
		"s", "05",

		"t", "pm",
		"T", "PM",
	}
	replacer := strings.NewReplacer(patterns...)
	str := replacer.Replace(format)
	return tm.Format(str)
}

// S时间_秒转时间文本 将秒数转换为中文可读的时间文本。
// 例如：3661 → "1时1分1秒"，31536000 → "1年"。
// 只显示非零的时间单位，所有单位为零时显示"0秒"。
//
// 参数:
//   - 秒: 总秒数
//
// 返回:
//   - string: 中文格式的时间文本（如 "1年2月3天4时5分6秒"）
func S时间_秒转时间文本(秒 int64) string {
	var 年, 月, 天, 时, 分, 剩余秒 int64
	const (
		每分秒数 = 60
		每时秒数 = 3600
		每天秒数 = 86400
		每月秒数 = 2592000
		每年秒数 = 31536000
	)

	年 = 秒 / 每年秒数
	秒 %= 每年秒数

	月 = 秒 / 每月秒数
	秒 %= 每月秒数

	天 = 秒 / 每天秒数
	秒 %= 每天秒数

	时 = 秒 / 每时秒数
	秒 %= 每时秒数

	分 = 秒 / 每分秒数
	剩余秒 = 秒 % 每分秒数

	result := ""
	if 年 > 0 {
		result += fmt.Sprintf("%d年", 年)
	}
	if 月 > 0 {
		result += fmt.Sprintf("%d月", 月)
	}
	if 天 > 0 {
		result += fmt.Sprintf("%d天", 天)
	}
	if 时 > 0 {
		result += fmt.Sprintf("%d时", 时)
	}
	if 分 > 0 {
		result += fmt.Sprintf("%d分", 分)
	}
	if 剩余秒 > 0 {
		result += fmt.Sprintf("%d秒", 剩余秒)
	}

	if result == "" {
		return "0秒"
	}

	return result
}
