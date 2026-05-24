package utils

import "sync/atomic"

// Y原子_递增 对 int64 变量进行原子递增 1 操作，线程安全。
// 等同于 atomic.AddInt64(&变量, 1)。
//
// 参数:
//   - 整数变量: 指向待递增变量的指针
//
// 返回:
//   - int64: 递增后的新值
//
// 示例:
//
//	var 计数器 int64 = 0
//	Y原子_递增(&计数器)  // 1
//	Y原子_递增(&计数器)  // 2
func Y原子_递增(整数变量 *int64) int64 {
	return atomic.AddInt64(整数变量, 1)
}

// Y原子_递减 对 int64 变量进行原子递减 1 操作，线程安全。
// 等同于 atomic.AddInt64(&变量, -1)。
//
// 参数:
//   - 整数变量: 指向待递减变量的指针
//
// 返回:
//   - int64: 递减后的新值
func Y原子_递减(整数变量 *int64) int64 {
	return atomic.AddInt64(整数变量, -1)
}
