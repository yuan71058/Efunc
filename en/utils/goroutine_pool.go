// 协程池工具
// 基于 panjf2000/ants/v2 库，提供 goroutine 池的创建、任务提交、状态监控等功能。
// 协程池可有效控制并发数量，避免 goroutine 泄漏和资源过度消耗。
package utils

import (
	"github.com/panjf2000/ants/v2"
)

// GoroutinePool_New 创建指定大小的 goroutine 池。
// 池满时新提交的任务会阻塞等待空闲 worker。
//
// 参数:
//   - size: worker 数量上限
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func GoroutinePool_New(size int) (*ants.Pool, error) {
	return ants.NewPool(size)
}

// GoroutinePool_NewWithOptions 创建带自定义选项的 goroutine 池。
//
// 参数:
//   - size: worker 数量上限
//   - options: ants.Option 列表，如 ants.WithPreAlloc(true)
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func GoroutinePool_NewWithOptions(size int, options ...ants.Option) (*ants.Pool, error) {
	return ants.NewPool(size, options...)
}

// GoroutinePool_Submit 向协程池提交一个任务。
// 如果池已满，会阻塞等待有空闲 worker。
//
// 参数:
//   - pool: 协程池实例
//   - task: 要执行的函数
//
// 返回:
//   - error: 提交失败时返回错误（如池已关闭）
func GoroutinePool_Submit(pool *ants.Pool, task func()) error {
	return pool.Submit(task)
}

// GoroutinePool_Running 获取当前正在执行任务的 worker 数量。
//
// 参数:
//   - pool: 协程池实例
//
// 返回:
//   - int: 运行中的 worker 数量
func GoroutinePool_Running(pool *ants.Pool) int {
	return pool.Running()
}

// GoroutinePool_Free 获取当前空闲的 worker 数量。
//
// 参数:
//   - pool: 协程池实例
//
// 返回:
//   - int: 空闲的 worker 数量
func GoroutinePool_Free(pool *ants.Pool) int {
	return pool.Free()
}

// GoroutinePool_Cap 获取协程池的容量（最大 worker 数）。
//
// 参数:
//   - pool: 协程池实例
//
// 返回:
//   - int: 池容量
func GoroutinePool_Cap(pool *ants.Pool) int {
	return pool.Cap()
}

// GoroutinePool_Waiting 获取正在等待执行的任务数量。
//
// 参数:
//   - pool: 协程池实例
//
// 返回:
//   - int: 等待中的任务数量
func GoroutinePool_Waiting(pool *ants.Pool) int {
	return pool.Waiting()
}

// GoroutinePool_Release 释放协程池，停止所有 worker 并清理资源。
// 释放后不能再提交任务。
//
// 参数:
//   - pool: 协程池实例
func GoroutinePool_Release(pool *ants.Pool) {
	pool.Release()
}

// GoroutinePool_Tune 动态调整协程池的容量。
// 新容量可以大于或小于当前容量。
//
// 参数:
//   - pool: 协程池实例
//   - newSize: 新的池容量
func GoroutinePool_Tune(pool *ants.Pool, newSize int) {
	pool.Tune(newSize)
}

// GoroutinePool_NewPreAlloc 创建预分配 worker 的协程池。
// 预分配模式下，worker 在创建池时即初始化，减少首次任务延迟。
//
// 参数:
//   - size: worker 数量上限
//
// 返回:
//   - *ants.Pool: 协程池实例
//   - error: 创建失败时返回错误
func GoroutinePool_NewPreAlloc(size int) (*ants.Pool, error) {
	return ants.NewPool(size, ants.WithPreAlloc(true))
}