package utils

import (
	"github.com/panjf2000/ants/v2"
)

// G协程池_创建 创建指定大小的 goroutine 池。
// 池满时新提交的任务会阻塞等待空闲 worker。
//
// 参数:
//   - 池大小: worker 数量上限
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func G协程池_创建(池大小 int) (*ants.Pool, error) {
	return ants.NewPool(池大小)
}

// G协程池_创建带选项 创建带自定义选项的 goroutine 池。
//
// 参数:
//   - 池大小: worker 数量上限
//   - 选项: ants.Option 列表，如 ants.WithPreAlloc(true)
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func G协程池_创建带选项(池大小 int, 选项 ...ants.Option) (*ants.Pool, error) {
	return ants.NewPool(池大小, 选项...)
}

// G协程池_提交任务 向协程池提交一个任务。
// 如果池已满，会阻塞等待有空闲 worker。
//
// 参数:
//   - 池: 协程池实例
//   - 任务: 要执行的函数
//
// 返回:
//   - error: 提交失败时返回错误（如池已关闭）
func G协程池_提交任务(池 *ants.Pool, 任务 func()) error {
	return 池.Submit(任务)
}

// G协程池_取运行中数量 获取当前正在执行任务的 worker 数量。
//
// 参数:
//   - 池: 协程池实例
//
// 返回:
//   - int: 运行中的 worker 数量
func G协程池_取运行中数量(池 *ants.Pool) int {
	return 池.Running()
}

// G协程池_取空闲数量 获取当前空闲的 worker 数量。
//
// 参数:
//   - 池: 协程池实例
//
// 返回:
//   - int: 空闲的 worker 数量
func G协程池_取空闲数量(池 *ants.Pool) int {
	return 池.Free()
}

// G协程池_取容量 获取协程池的容量（最大 worker 数）。
//
// 参数:
//   - 池: 协程池实例
//
// 返回:
//   - int: 池容量
func G协程池_取容量(池 *ants.Pool) int {
	return 池.Cap()
}

// G协程池_取等待数量 获取正在等待执行的 任务 数量。
//
// 参数:
//   - 池: 协程池实例
//
// 返回:
//   - int: 等待中的 任务 数量
func G协程池_取等待数量(池 *ants.Pool) int {
	return 池.Waiting()
}

// G协程池_释放 释放协程池，停止所有 worker 并清理资源。
// 释放后不能再提交任务。
//
// 参数:
//   - 池: 协程池实例
func G协程池_释放(池 *ants.Pool) {
	池.Release()
}

// G协程池_调整大小 动态调整协程池的容量。
// 新容量可以大于或小于当前容量。
//
// 参数:
//   - 池: 协程池实例
//   - 新大小: 新的池容量
func G协程池_调整大小(池 *ants.Pool, 新大小 int) {
	池.Tune(新大小)
}

// G协程池_预分配 创建预分配 worker 的协程池。
// 预分配模式下，worker 在创建池时即初始化，减少首次任务延迟。
//
// 参数:
//   - 池大小: worker 数量上限
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func G协程池_预分配(池大小 int) (*ants.Pool, error) {
	return ants.NewPool(池大小, ants.WithPreAlloc(true))
}
