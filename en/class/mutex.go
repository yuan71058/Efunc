package class

import "sync"

// Mutex 互斥锁封装，用于保护临界区代码。
// 基于 sync.Mutex 实现，确保同一时刻只有一个 goroutine 进入临界区。
//
// 使用示例:
//
//	lock := &Mutex{}
//	lock.Lock()
//	// 临界区代码，同一时刻只有一个 goroutine 执行
//	lock.Unlock()
type Mutex struct {
	lock sync.Mutex
}

// Lock 获取互斥锁，进入临界区。
// 如果锁已被其他 goroutine 持有，则阻塞等待。
// 必须与 Unlock 配对使用，否则会导致死锁。
func (l *Mutex) Lock() {
	l.lock.Lock()
}

// Unlock 释放互斥锁，退出临界区。
// 必须与 Lock 配对使用。
func (l *Mutex) Unlock() {
	l.lock.Unlock()
}

// TryLock 尝试获取互斥锁而不阻塞。
// 如果锁当前可用则获取并返回 true，否则立即返回 false。
//
// 返回:
//   - bool: true 表示成功获取锁，false 表示锁已被占用
func (l *Mutex) TryLock() bool {
	return l.lock.TryLock()
}