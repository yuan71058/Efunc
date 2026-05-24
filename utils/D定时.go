package utils

import (
	"github.com/robfig/cron/v3"
)

// D定时_创建 创建一个新的 cron 定时任务调度器。
// 支持标准 cron 表达式和秒级 cron 表达式。
//
// 返回:
//   - *cron.Cron: cron 调度器实例
func D定时_创建() *cron.Cron {
	return cron.New(cron.WithSeconds())
}

// D定时_添加任务 向调度器添加定时任务。
// 表达式格式（秒级）：秒 分 时 日 月 周
// 例如 "*/5 * * * * *" 表示每5秒执行一次
// "0 30 9 * * *" 表示每天9:30执行
//
// 参数:
//   - 调度器: cron 调度器实例
//   - 表达式: cron 表达式
//   - 任务: 要执行的函数
//
// 返回:
//   - cron.EntryID: 任务 ID，可用于后续移除
//   - error: 表达式无效时返回错误
func D定时_添加任务(调度器 *cron.Cron, 表达式 string, 任务 func()) (cron.EntryID, error) {
	return 调度器.AddFunc(表达式, 任务)
}

// D定时_移除任务 从调度器中移除指定任务。
//
// 参数:
//   - 调度器: cron 调度器实例
//   - 任务ID: 要移除的任务 ID
func D定时_移除任务(调度器 *cron.Cron, 任务ID cron.EntryID) {
	调度器.Remove(任务ID)
}

// D定时_启动 启动调度器，开始执行所有已添加的定时任务。
// 调度器会在后台 goroutine 中运行。
//
// 参数:
//   - 调度器: cron 调度器实例
func D定时_启动(调度器 *cron.Cron) {
	调度器.Start()
}

// D定时_停止 停止调度器，不再执行定时任务。
//
// 参数:
//   - 调度器: cron 调度器实例
func D定时_停止(调度器 *cron.Cron) {
	调度器.Stop()
}

// D定时_取任务列表 获取调度器中所有已注册的任务。
//
// 参数:
//   - 调度器: cron 调度器实例
//
// 返回:
//   - []cron.Entry: 任务条目列表
func D定时_取任务列表(调度器 *cron.Cron) []cron.Entry {
	return 调度器.Entries()
}

// D定时_简单执行 按指定 cron 表达式周期性执行任务。
// 创建调度器、添加任务、启动一条龙完成。
//
// 参数:
//   - 表达式: cron 表达式（秒级），如 "*/10 * * * * *"
//   - 任务: 要执行的函数
//
// 返回:
//   - *cron.Cron: 调度器实例，可用于后续停止
//   - error: 表达式无效时返回错误
func D定时_简单执行(表达式 string, 任务 func()) (*cron.Cron, error) {
	调度器 := D定时_创建()
	_, 错误 := D定时_添加任务(调度器, 表达式, 任务)
	if 错误 != nil {
		return nil, 错误
	}
	D定时_启动(调度器)
	return 调度器, nil
}
