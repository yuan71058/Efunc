package utils

import (
	"fmt"
	"strings"
	"time"
)

const base_format = "2006-01-02 15:04:05"

// S时间_文本到时间戳
// @时间文本  2006-01-02 15:04:05
func S时间_文本到时间戳(时间文本 string) int {

	formatTime, err := time.ParseInLocation("2006-01-02 15:04:05", 时间文本, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	if err == nil {
		return int(formatTime.Unix())
	}
	return 0

}

// S时间_取现行时间戳13
func S时间_取现行时间戳13() int64 {
	return time.Now().UnixNano() / 1e6
}

// S时间_取现行时间戳
// 返回十位时间戳
func S时间_取现行时间戳() int64 {
	return time.Now().Unix()
}

// 调用格式： 〈日期时间型〉 取现行时间 （） - 系统核心支持库->时间操作
// 英文名称：now
// 返回当前系统日期及时间。本命令为初级命令。
//
// 操作系统需求： Windows、Linux
func S时间_取现行时间() string {
	//获取当前时间
	t := time.Now()
	//2019-02-21 17:20:57.0764497 +0800 CST m=+0.018555201
	str_time := t.Format(base_format)
	//2019-02-21 17:20:57
	return str_time
}

func S时间_时间戳到时间(时间戳 int64) string {
	return time.Unix(时间戳, 0).Format(base_format)
}
func S时间_时间戳13到时间(时间戳 int64) string {
	return time.Unix(时间戳/1000, 0).Format(base_format)
}
func S时间_时间到时间戳(时间 string) int64 {
	layout := "2006-01-02 15:04:05"

	// 解析字符串为 time.Time 对象
	t, err := time.Parse(layout, 时间)
	if err != nil {
		return 0
	}
	// 获取 Unix 时间戳（秒）
	timestampInSeconds := t.Unix()
	return timestampInSeconds
}

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

		"h", "3", //12小时制
		"H", "15", //24小时制

		"i", "04",
		"s", "05",

		"t", "pm",
		"T", "PM",
	}
	replacer := strings.NewReplacer(patterns...)
	str := replacer.Replace(format)
	return tm.Format(str)
}

// 秒转时间文本 将给定的秒数转换为 年、月、日、时、分、秒 的文本表示，只显示非零项
func S时间_秒转时间文本(秒 int64) string {
	var 年, 月, 天, 时, 分, 剩余秒 int64
	const (
		每分秒数 = 60
		每时秒数 = 3600
		每天秒数 = 86400
		每月秒数 = 2592000  // 按30天计算
		每年秒数 = 31536000 // 按365天计算（即12个月）
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

	// 构建结果字符串，只包含非零项
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

	// 如果所有单位都是0，则显示0秒
	if result == "" {
		return "0秒"
	}

	return result
}
