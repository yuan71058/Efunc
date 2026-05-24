package class

import "sync"

// L_读写锁 读写锁封装，基于 sync.RWMutex 实现。
// 读共享：多个 goroutine 可以同时获取读锁，适合读多写少的场景。
// 写独享：写锁会阻塞所有读锁和其他写锁，保证数据一致性。
//
// 使用示例:
//
//	lock := &L_读写锁{}
//	lock.K开始读()
//	// 读取共享数据...
//	lock.J结束读()
//	lock.K开始写()
//	// 修改共享数据...
//	lock.J结束写()
type L_读写锁 struct {
	lock sync.RWMutex
}

// K开始读 获取读锁。
// 多个 goroutine 可以同时持有读锁，但写锁被持有时会阻塞。
func (l *L_读写锁) K开始读() {
	l.lock.RLock()
}

// J结束读 释放读锁。
// 必须与 K开始读 配对使用，否则会导致死锁。
func (l *L_读写锁) J结束读() {
	l.lock.RUnlock()
}

// K开始写 获取写锁。
// 写锁是独占的，获取写锁时会阻塞所有读锁和其他写锁。
func (l *L_读写锁) K开始写() {
	l.lock.Lock()
}

// J结束写 释放写锁。
// 必须与 K开始写 配对使用，否则会导致死锁。
func (l *L_读写锁) J结束写() {
	l.lock.Unlock()
}
