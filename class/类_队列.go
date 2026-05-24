package class

import (
	"container/list"
	"fmt"
	"sync"
)

// L_队列 先入先出（FIFO）队列实现，线程安全。
// 基于 container/list 双向链表，使用互斥锁保证并发安全。
// 元素类型为 interface{}，可存储任意类型的值。
//
// 使用示例:
//
//	q := &L_队列{}
//	q.Init()
//	q.J加入队列("hello")
//	val, ok := q.T弹出队列()
type L_队列 struct {
	l      sync.Mutex
	data   list.List
	isInit bool
}

// Init 初始化队列，清空所有已有数据。
func (j *L_队列) Init() {
	j.data.Init()
}

// J加入队列 将元素加入队列尾部，返回当前队列长度。
//
// 参数:
//   - v: 要加入队列的元素，类型为 interface{}
//
// 返回:
//   - int: 加入后队列中的元素数量
func (q *L_队列) J加入队列(v interface{}) int {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.PushFront(v)
	return q.data.Len()
}

// T弹出队列 从队列头部弹出一个元素（先进先出）。
//
// 返回:
//   - interface{}: 弹出的元素值；队列为空时返回 nil
//   - bool: true 表示成功弹出，false 表示队列为空
func (q *L_队列) T弹出队列() (interface{}, bool) {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if nil == iter {
		return nil, false
	}
	v := iter.Value
	q.data.Remove(iter)
	return v, true
}

// T弹出队列文本 从队列头部弹出一个字符串类型的元素。
// 如果队列为空或元素不是字符串类型，返回 false。
//
// 参数:
//   - 值: 用于接收弹出字符串的指针
//
// 返回:
//   - bool: true 表示成功弹出字符串，false 表示队列为空或类型不匹配
func (q *L_队列) T弹出队列文本(值 *string) bool {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if nil == iter {
		return false
	}
	v := iter.Value
	局_临时, ok := v.(string)
	if !ok {
		return false
	}
	q.data.Remove(iter)
	*值 = 局_临时
	return true
}

// Q取队列长度 获取队列中当前元素的数量。
//
// 返回:
//   - int: 队列长度
func (q *L_队列) Q取队列长度() int {
	return q.data.Len()
}

// Q清空队列 清空队列中的所有元素。
//
// 返回:
//   - interface{}: 内部 list.Init() 的返回值
func (q *L_队列) Q清空队列() interface{} {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.Init()
}

// Dump 打印队列中所有元素的内容，用于调试。
// 从队首到队尾依次输出，格式为 "item: <value>"。
func (q *L_队列) Dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}

// T弹出队列整数 从队列头部弹出一个整数类型的元素。
// 如果队列为空或元素不是 int 类型，返回 false。
//
// 参数:
//   - 值: 用于接收弹出整数的指针
//
// 返回:
//   - bool: true 表示成功弹出整数，false 表示队列为空或类型不匹配
func (q *L_队列) T弹出队列整数(值 *int) bool {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if nil == iter {
		return false
	}
	v := iter.Value
	局_临时, ok := v.(int)
	if !ok {
		return false
	}
	q.data.Remove(iter)
	*值 = 局_临时
	return true
}
