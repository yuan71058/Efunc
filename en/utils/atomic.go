// 原子操作工具
// 基于 sync/atomic 包，提供线程安全的递增、递减等原子操作。
package utils

import "sync/atomic"

// Atomic_Increment 对 int64 变量进行原子递增 1 操作，线程安全。
// 等同于 atomic.AddInt64(&变量, 1)。
//
// 参数:
//   - p: 指向待递增变量的指针
//
// 返回:
//   - int64: 递增后的新值
//
// 示例:
//
//	var counter int64 = 0
//	Atomic_Increment(&counter)  // 1
//	Atomic_Increment(&counter)  // 2
func Atomic_Increment(p *int64) int64 {
	return atomic.AddInt64(p, 1)
}

// Atomic_Decrement 对 int64 变量进行原子递减 1 操作，线程安全。
// 等同于 atomic.AddInt64(&变量, -1)。
//
// 参数:
//   - p: 指向待递减变量的指针
//
// 返回:
//   - int64: 递减后的新值
func Atomic_Decrement(p *int64) int64 {
	return atomic.AddInt64(p, -1)
}