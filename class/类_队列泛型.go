package class

import (
	"container/list"
	"fmt"
	"sync"
)

// L_队列泛型 线程安全的泛型队列实现。
// 基于 container/list 双向链表，使用互斥锁保证并发安全。
// 支持任意类型的元素，遵循先进先出（FIFO）原则。
//
// 使用示例:
//
//	q := &L_队列泛型[int]{}
//	q.Init()
//	q.J加入队列(42)
//	val, ok := q.T弹出队列()
type L_队列泛型[T any] struct {
	l      sync.Mutex
	data   list.List
	isInit bool
}

// Init 初始化队列，清空所有已有数据。
// 在首次使用前调用，确保队列处于干净状态。
func (q *L_队列泛型[T]) Init() {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.Init()
}

// J加入队列 将元素加入队列尾部，返回当前队列长度。
//
// 参数:
//   - v: 要加入队列的元素
//
// 返回:
//   - int: 加入后队列中的元素数量
func (q *L_队列泛型[T]) J加入队列(v T) int {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.PushFront(v)
	return q.data.Len()
}

// T弹出队列 从队列头部弹出一个元素（先进先出）。
//
// 返回:
//   - T: 弹出的元素值；队列为空时返回类型零值
//   - bool: true 表示成功弹出，false 表示队列为空
func (q *L_队列泛型[T]) T弹出队列() (T, bool) {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if iter == nil {
		var zero T
		return zero, false
	}
	v := iter.Value.(T)
	q.data.Remove(iter)
	return v, true
}

// Q取队列长度 获取队列中当前元素的数量。
//
// 返回:
//   - int: 队列长度
func (q *L_队列泛型[T]) Q取队列长度() int {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.Len()
}

// Q清空队列 清空队列中的所有元素。
func (q *L_队列泛型[T]) Q清空队列() {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.Init()
}

// Dump 打印队列中所有元素的内容，用于调试。
// 从队首到队尾依次输出，格式为 "item: <value>"。
func (q *L_队列泛型[T]) Dump() {
	q.l.Lock()
	defer q.l.Unlock()
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}
