package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
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

// S时间_取网络时间 通过 NTP 协议从指定 NTP 服务器获取网络标准时间。
// 使用 SNTP 协议（RFC 4330）与 NTP 服务器通信，获取精确的网络时间。
//
// 参数:
//   - 服务器: NTP 服务器地址，如 "ntp.aliyun.com"、"time.windows.com"；空字符串使用默认服务器
//
// 返回:
//   - time.Time: 网络标准时间；获取失败返回零值
func S时间_取网络时间(服务器 string) time.Time {
	if 服务器 == "" {
		服务器 = "ntp.aliyun.com"
	}
	conn, err := net.DialTimeout("udp", 服务器+":123", 5*time.Second)
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

// S时间_取网络时间戳 通过 NTP 协议获取网络标准时间的 10 位秒级时间戳。
//
// 参数:
//   - 服务器: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - int64: 10 位秒级时间戳；获取失败返回 0
func S时间_取网络时间戳(服务器 string) int64 {
	t := S时间_取网络时间(服务器)
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}

// S时间_取网络时间文本 通过 NTP 协议获取网络标准时间并格式化为文本。
//
// 参数:
//   - 服务器: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - string: 格式为 "2006-01-02 15:04:05" 的网络时间字符串；获取失败返回空字符串
func S时间_取网络时间文本(服务器 string) string {
	t := S时间_取网络时间(服务器)
	if t.IsZero() {
		return ""
	}
	return t.Format(base_format)
}

// S时间_取HTTP网络时间 通过 HTTP 请求从世界时间 API 获取网络时间。
// 适用于 NTP 端口被封锁但 HTTP 可用的网络环境。
//
// 返回:
//   - time.Time: 网络标准时间；获取失败返回零值
func S时间_取HTTP网络时间() time.Time {
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

// S时间_取本地与网络时差 获取本地时间与网络时间的差值（秒）。
// 正值表示本地时间快于网络时间，负值表示本地时间慢于网络时间。
//
// 参数:
//   - 服务器: NTP 服务器地址；空字符串使用默认服务器
//
// 返回:
//   - int64: 时差（秒）；获取失败返回 0
func S时间_取本地与网络时差(服务器 string) int64 {
	网络时间 := S时间_取网络时间(服务器)
	if 网络时间.IsZero() {
		return 0
	}
	return time.Now().Unix() - 网络时间.Unix()
}

// S时间_取日期 获取当前日期，格式为 "2006-01-02"。
//
// 返回:
//   - string: 当前日期字符串
func S时间_取日期() string {
	return time.Now().Format("2006-01-02")
}

// S时间_取时间 获取当前时间，格式为 "15:04:05"。
//
// 返回:
//   - string: 当前时间字符串
func S时间_取时间() string {
	return time.Now().Format("15:04:05")
}

// S时间_取星期 获取当前日期是星期几的中文名称。
//
// 返回:
//   - string: 星期名称（星期日~星期六）
func S时间_取星期() string {
	星期 := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	return 星期[time.Now().Weekday()]
}

// S时间_计算时间差 计算两个时间戳之间的时间差，返回中文可读文本。
//
// 参数:
//   - 开始时间戳: 10 位秒级时间戳
//   - 结束时间戳: 10 位秒级时间戳
//
// 返回:
//   - string: 中文格式的时间差文本（如 "1时2分3秒"）
func S时间_计算时间差(开始时间戳, 结束时间戳 int64) string {
	差值 := 结束时间戳 - 开始时间戳
	if 差值 < 0 {
		差值 = -差值
	}
	return S时间_秒转时间文本(差值)
}

// S时间_是否闰年 判断指定年份是否为闰年。
//
// 参数:
//   - 年: 年份，如 2024
//
// 返回:
//   - bool: 闰年返回 true，平年返回 false
func S时间_是否闰年(年 int) bool {
	return 年%4 == 0 && (年%100 != 0 || 年%400 == 0)
}

// S时间_取月份天数 获取指定年份和月份的天数。
//
// 参数:
//   - 年: 年份，如 2024
//   - 月: 月份（1-12）
//
// 返回:
//   - int: 该月天数；月份无效返回 0
func S时间_取月份天数(年, 月 int) int {
	if 月 < 1 || 月 > 12 {
		return 0
	}
	return time.Date(年, time.Month(月+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
