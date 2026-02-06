package class

import (
	"container/list"
	"fmt"
	"sync"
)

// 泛型队列，线程安全
type L_队列泛型[T any] struct {
	l      sync.Mutex
	data   list.List
	isInit bool
}

// 初始化队列
func (q *L_队列泛型[T]) Init() {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.Init()
}

// 加入队列，返回当前队列长度
func (q *L_队列泛型[T]) J加入队列(v T) int {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.PushFront(v)
	return q.data.Len()
}

// 弹出队列元素，返回值和是否成功
func (q *L_队列泛型[T]) T弹出队列() (T, bool) {
	q.l.Lock()
	defer q.l.Unlock()
	iter := q.data.Back()
	if iter == nil {
		var zero T // 零值
		return zero, false
	}
	v := iter.Value.(T)
	q.data.Remove(iter)
	return v, true
}

// 获取队列长度
func (q *L_队列泛型[T]) Q取队列长度() int {
	q.l.Lock()
	defer q.l.Unlock()
	return q.data.Len()
}

// 清空队列
func (q *L_队列泛型[T]) Q清空队列() {
	q.l.Lock()
	defer q.l.Unlock()
	q.data.Init()
}

// 打印队列内容（调试用）
func (q *L_队列泛型[T]) Dump() {
	q.l.Lock()
	defer q.l.Unlock()
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}
