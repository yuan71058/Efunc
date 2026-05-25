package class

import "sync"

// RWMutex 读写锁封装，基于 sync.RWMutex 实现。
// 读共享：多个 goroutine 可以同时获取读锁，适合读多写少的场景。
// 写独享：写锁会阻塞所有读锁和其他写锁，保证数据一致性。
//
// 使用示例:
//
//	lock := &RWMutex{}
//	lock.RLock()
//	// 读取共享数据...
//	lock.RUnlock()
//	lock.Lock()
//	// 修改共享数据...
//	lock.Unlock()
type RWMutex struct {
	lock sync.RWMutex
}

// RLock 获取读锁。
// 多个 goroutine 可以同时持有读锁，但写锁被持有时会阻塞。
func (l *RWMutex) RLock() {
	l.lock.RLock()
}

// RUnlock 释放读锁。
// 必须与 RLock 配对使用，否则会导致死锁。
func (l *RWMutex) RUnlock() {
	l.lock.RUnlock()
}

// Lock 获取写锁。
// 写锁是独占的，获取写锁时会阻塞所有读锁和其他写锁。
func (l *RWMutex) Lock() {
	l.lock.Lock()
}

// Unlock 释放写锁。
// 必须与 Lock 配对使用，否则会导致死锁。
func (l *RWMutex) Unlock() {
	l.lock.Unlock()
}