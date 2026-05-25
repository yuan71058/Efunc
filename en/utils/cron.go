// 定时任务调度工具
// 基于 robfig/cron/v3 库，支持秒级 cron 表达式的定时任务调度。
// 表达式格式（秒级）：秒 分 时 日 月 周
// 例如 "*/5 * * * * *" 表示每5秒执行一次，"0 30 9 * * *" 表示每天9:30执行。
package utils

import (
	"github.com/robfig/cron/v3"
)

// Cron_New 创建一个新的 cron 定时任务调度器。
// 支持标准 cron 表达式和秒级 cron 表达式。
//
// 返回:
//   - *cron.Cron: cron 调度器实例
func Cron_New() *cron.Cron {
	return cron.New(cron.WithSeconds())
}

// Cron_AddFunc 向调度器添加定时任务。
// 表达式格式（秒级）：秒 分 时 日 月 周
// 例如 "*/5 * * * * *" 表示每5秒执行一次
// "0 30 9 * * *" 表示每天9:30执行
//
// 参数:
//   - c: cron 调度器实例
//   - spec: cron 表达式
//   - cmd: 要执行的函数
//
// 返回:
//   - cron.EntryID: 任务 ID，可用于后续移除
//   - error: 表达式无效时返回错误
func Cron_AddFunc(c *cron.Cron, spec string, cmd func()) (cron.EntryID, error) {
	return c.AddFunc(spec, cmd)
}

// Cron_Remove 从调度器中移除指定任务。
//
// 参数:
//   - c: cron 调度器实例
//   - id: 要移除的任务 ID
func Cron_Remove(c *cron.Cron, id cron.EntryID) {
	c.Remove(id)
}

// Cron_Start 启动调度器，开始执行所有已添加的定时任务。
// 调度器会在后台 goroutine 中运行。
//
// 参数:
//   - c: cron 调度器实例
func Cron_Start(c *cron.Cron) {
	c.Start()
}

// Cron_Stop 停止调度器，不再执行定时任务。
//
// 参数:
//   - c: cron 调度器实例
func Cron_Stop(c *cron.Cron) {
	c.Stop()
}

// Cron_Entries 获取调度器中所有已注册的任务。
//
// 参数:
//   - c: cron 调度器实例
//
// 返回:
//   - []cron.Entry: 任务条目列表
func Cron_Entries(c *cron.Cron) []cron.Entry {
	return c.Entries()
}

// Cron_Run 按指定 cron 表达式周期性执行任务。
// 创建调度器、添加任务、启动一条龙完成。
//
// 参数:
//   - spec: cron 表达式（秒级），如 "*/10 * * * * *"
//   - cmd: 要执行的函数
//
// 返回:
//   - *cron.Cron: 调度器实例，可用于后续停止
//   - error: 表达式无效时返回错误
func Cron_Run(spec string, cmd func()) (*cron.Cron, error) {
	c := Cron_New()
	_, err := Cron_AddFunc(c, spec, cmd)
	if err != nil {
		return nil, err
	}
	Cron_Start(c)
	return c, nil
}