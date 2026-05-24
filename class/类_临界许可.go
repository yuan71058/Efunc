package class

import "sync"

// L_临界许可 互斥锁封装，用于保护临界区代码。
// 基于 sync.Mutex 实现，确保同一时刻只有一个 goroutine 进入临界区。
//
// 使用示例:
//
//	lock := &L_临界许可{}
//	lock.J进入许可区()
//	// 临界区代码，同一时刻只有一个 goroutine 执行
//	lock.T退出许可区()
type L_临界许可 struct {
	lock sync.Mutex
}

// J进入许可区 获取互斥锁，进入临界区。
// 如果锁已被其他 goroutine 持有，则阻塞等待。
// 必须与 T退出许可区 配对使用，否则会导致死锁。
func (l *L_临界许可) J进入许可区() {
	l.lock.Lock()
}

// T退出许可区 释放互斥锁，退出临界区。
// 必须与 J进入许可区 配对使用。
func (l *L_临界许可) T退出许可区() {
	l.lock.Unlock()
}

// C尝试进入 尝试获取互斥锁而不阻塞。
// 如果锁当前可用则获取并返回 true，否则立即返回 false。
//
// 返回:
//   - bool: true 表示成功获取锁，false 表示锁已被占用
func (l *L_临界许可) C尝试进入() bool {
	return l.lock.TryLock()
}
