package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// 默认的时间格式模板，对应 Go 的 "2006-01-02 15:04:05"
const defaultTimeFormat = "2006-01-02 15:04:05"

// Time_TextToTimestamp 将时间文本转换为 10 位 Unix 时间戳（秒）。
// 使用本地时区解析，格式必须为 "2006-01-02 15:04:05"。
//
// 参数:
//   - timeStr: 格式为 "2006-01-02 15:04:05" 的时间字符串
//
// 返回:
//   - int: 10 位 Unix 时间戳；解析失败返回 0
func Time_TextToTimestamp(timeStr string) int {
	formatTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err == nil {
		return int(formatTime.Unix())
	}
	return 0
}

// Time_NowMillisecond 获取当前时间的 13 位毫秒级时间戳。
//
// 返回:
//   - int64: 13 位毫秒级时间戳
func Time_NowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// Time_NowTimestamp 获取当前时间的 10 位秒级时间戳。
//
// 返回:
//   - int64: 10 位秒级时间戳
func Time_NowTimestamp() int64 {
	return time.Now().Unix()
}

// Time_Now 获取当前系统日期和时间，格式为 "2006-01-02 15:04:05"。
//
// 返回:
//   - string: 格式化的当前时间字符串
func Time_Now() string {
	t := time.Now()
	strTime := t.Format(defaultTimeFormat)
	return strTime
}

// Time_TimestampToTime 将 10 位秒级时间戳转换为格式化的时间字符串。
//
// 参数:
//   - timestamp: 10 位秒级时间戳
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的时间字符串
func Time_TimestampToTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(defaultTimeFormat)
}

// Time_MillisecondToTime 将 13 位毫秒级时间戳转换为格式化的时间字符串。
//
// 参数:
//   - timestamp: 13 位毫秒级时间戳
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的时间字符串
func Time_MillisecondToTime(timestamp int64) string {
	return time.Unix(timestamp/1000, 0).Format(defaultTimeFormat)
}

// Time_TimeToTimestamp 将格式化的时间字符串转换为 10 位秒级时间戳。
// 使用 UTC 时区解析，与 Time_TextToTimestamp（本地时区）可能有时区差异。
//
// 参数:
//   - timeStr: 格式为 "2006-01-02 15:04:05" 的时间字符串
//
// 返回:
//   - int64: 10 位秒级时间戳；解析失败返回 0
func Time_TimeToTimestamp(timeStr string) int64 {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0
	}
	timestampInSeconds := t.Unix()
	return timestampInSeconds
}

// Time_FormatTimestamp 使用自定义格式字符串格式化时间戳。
// 支持的占位符：y/Y=年, m/M=月, d/D=日, h=12小时制, H=24小时制, i=分, s=秒, t=am/pm, T=AM/PM。
// 时间戳为 0 时使用当前时间。
//
// 参数:
//   - format: 自定义格式字符串（如 "Y-m-d H:i:s"）
//   - timestamp: 10 位秒级时间戳，0 表示当前时间
//
// 返回:
//   - string: 格式化后的时间字符串
func Time_FormatTimestamp(format string, timestamp int64) string {
	var tm time.Time

	if timestamp == 0 {
		tm = time.Now()
	} else {
		tm = time.Unix(timestamp, 0)
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

// Time_SecondsToText 将秒数转换为中文可读的时间文本。
// 例如：3661 → "1时1分1秒"，31536000 → "1年"。
// 只显示非零的时间单位，所有单位为零时显示"0秒"。
//
// 参数:
//   - seconds: 总秒数
//
// 返回:
//   - string: 中文格式的时间文本（如 "1年2月3天4时5分6秒"）
func Time_SecondsToText(seconds int64) string {
	var year, month, day, hour, minute, remainSeconds int64
	const (
		secondsPerMinute = 60
		secondsPerHour   = 3600
		secondsPerDay    = 86400
		secondsPerMonth  = 2592000
		secondsPerYear   = 31536000
	)

	year = seconds / secondsPerYear
	seconds %= secondsPerYear

	month = seconds / secondsPerMonth
	seconds %= secondsPerMonth

	day = seconds / secondsPerDay
	seconds %= secondsPerDay

	hour = seconds / secondsPerHour
	seconds %= secondsPerHour

	minute = seconds / secondsPerMinute
	remainSeconds = seconds % secondsPerMinute

	result := ""
	if year > 0 {
		result += fmt.Sprintf("%d年", year)
	}
	if month > 0 {
		result += fmt.Sprintf("%d月", month)
	}
	if day > 0 {
		result += fmt.Sprintf("%d天", day)
	}
	if hour > 0 {
		result += fmt.Sprintf("%d时", hour)
	}
	if minute > 0 {
		result += fmt.Sprintf("%d分", minute)
	}
	if remainSeconds > 0 {
		result += fmt.Sprintf("%d秒", remainSeconds)
	}

	if result == "" {
		return "0秒"
	}

	return result
}

// Time_GetNetworkTime 通过 NTP 协议从指定 NTP 服务器获取网络标准时间。
// 使用 SNTP 协议（RFC 4330）与 NTP 服务器通信，获取精确的网络时间。
//
// 参数:
//   - server: NTP 服务器地址，如 "ntp.aliyun.com"、"time.windows.com"；空字符串使用默认服务器
//
// 返回:
//   - time.Time: 网络标准时间；获取失败返回零值
func Time_GetNetworkTime(server string) time.Time {
	if server == "" {
		server = "ntp.aliyun.com"
	}
	conn, err := net.DialTimeout("udp", server+":123", 5*time.Second)
	if err != nil {
		return time.Time{}
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	req := make([]byte, 48)
	req[0] = 0x1b
	_, err = conn.Write(req)
	if err != nil {
		return time.Time{}
	}

	resp := make([]byte, 48)
	_, err = conn.Read(resp)
	if err != nil {
		return time.Time{}
	}

	sec := uint64(resp[40])<<24 | uint64(resp[41])<<16 | uint64(resp[42])<<8 | uint64(resp[43])
	frac := uint64(resp[44])<<24 | uint64(resp[45])<<16 | uint64(resp[46])<<8 | uint64(resp[47])

	nsec := frac * 1e9 / (1 << 32)
	unixSec := int64(sec) - 2208988800
	return time.Unix(unixSec, int64(nsec))
}

// Time_GetNetworkTimestamp 通过 NTP 协议获取网络标准时间的 10 位秒级时间戳。
//
// 参数:
//   - server: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - int64: 10 位秒级时间戳；获取失败返回 0
func Time_GetNetworkTimestamp(server string) int64 {
	t := Time_GetNetworkTime(server)
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}

// Time_GetNetworkTimeText 通过 NTP 协议获取网络标准时间并格式化为文本。
//
// 参数:
//   - server: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的网络时间字符串；获取失败返回空字符串
func Time_GetNetworkTimeText(server string) string {
	t := Time_GetNetworkTime(server)
	if t.IsZero() {
		return ""
	}
	return t.Format(defaultTimeFormat)
}

// Time_GetHTTPNetworkTime 通过 HTTP 请求从世界时间 API 获取网络时间。
// 适用于 NTP 端口被封锁但 HTTP 可用的网络环境。
//
// 返回:
//   - time.Time: 网络标准时间；获取失败返回零值
func Time_GetHTTPNetworkTime() time.Time {
	type timeAPIResp struct {
		Unixtime int64  `json:"unixtime"`
		Datetime string `json:"datetime"`
	}
	resp, err := http.DefaultClient.Get("http://worldtimeapi.org/api/timezone/Asia/Shanghai")
	if err != nil {
		return time.Time{}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return time.Time{}
	}
	var result timeAPIResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return time.Time{}
	}
	if result.Unixtime > 0 {
		return time.Unix(result.Unixtime, 0)
	}
	return time.Time{}
}

// Time_GetLocalNetworkDiff 获取本地时间与网络时间的差值（秒）。
// 正值表示本地时间快于网络时间，负值表示本地时间慢于网络时间。
//
// 参数:
//   - server: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - int64: 时差（秒）；获取失败返回 0
func Time_GetLocalNetworkDiff(server string) int64 {
	netTime := Time_GetNetworkTime(server)
	if netTime.IsZero() {
		return 0
	}
	return time.Now().Unix() - netTime.Unix()
}

// Time_GetDate 获取当前日期，格式为 "2006-01-02"。
//
// 返回:
//   - string: 当前日期字符串
func Time_GetDate() string {
	return time.Now().Format("2006-01-02")
}

// Time_GetTime 获取当前时间，格式为 "15:04:05"。
//
// 返回:
//   - string: 当前时间字符串
func Time_GetTime() string {
	return time.Now().Format("15:04:05")
}

// Time_GetWeekday 获取当前日期是星期几的中文名称。
//
// 返回:
//   - string: 星期名称（星期日~星期六）
func Time_GetWeekday() string {
	weekdays := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	return weekdays[time.Now().Weekday()]
}

// Time_CalcDiff 计算两个时间戳之间的时间差，返回中文可读文本。
//
// 参数:
//   - startTimestamp: 10 位秒级时间戳
//   - endTimestamp: 10 位秒级时间戳
//
// 返回:
//   - string: 中文格式的时间差文本（如 "1时2分3秒"）
func Time_CalcDiff(startTimestamp, endTimestamp int64) string {
	diff := endTimestamp - startTimestamp
	if diff < 0 {
		diff = -diff
	}
	return Time_SecondsToText(diff)
}

// Time_IsLeapYear 判断指定年份是否为闰年。
//
// 参数:
//   - year: 年份，如 2024
//
// 返回:
//   - bool: 闰年返回 true，平年返回 false
func Time_IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// Time_GetMonthDays 获取指定年份和月份的天数。
//
// 参数:
//   - year: 年份，如 2024
//   - month: 月份（1-12）
//
// 返回:
//   - int: 该月天数；月份无效返回 0
func Time_GetMonthDays(year, month int) int {
	if month < 1 || month > 12 {
		return 0
	}
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}