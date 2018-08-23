package queue

import (
	"errors"
)

type CQueue struct {
	ele  []interface{} // 存储size-1个
	size int64         // 容量
	head int64         // 头元素位置
	tail int64         // 队尾的下一个位置
}

func NewCQueue(size int64) (*CQueue, error) {
	if size <= 0 {
		return nil, errors.New("error size")
	}
	return &CQueue{
		ele:  make([]interface{}, size),
		size: size,
	}, nil
}

func (q CQueue) IsEmpty() bool {
	return q.head == q.tail && q.ele[q.tail] == nil
}

func (q CQueue) IsFull() bool {
	if (q.tail+1)%q.size == q.head {
		return true
	}
	return false
}

func (q CQueue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	if q.tail > q.head {
		return int(q.tail - q.head)
	} else {
		return int(q.size - (q.head - q.tail))
	}
}

// 入队
func (q *CQueue) Offer(ele interface{}) (bool, error) {
	if q.IsFull() {
		return false, errors.New("queue is full")
	}
	if q.IsEmpty() { // 插入第一个 head tail 都是0
		q.ele[q.tail] = ele
	}
	q.ele[q.tail] = ele
	q.tail = (q.tail + 1) % q.size
	return true, nil
}

// 出队并删除
func (q *CQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.ele[q.head]
	q.ele[q.head] = nil
	q.head = (q.head + 1) % q.size
	return ret
}

func (q CQueue) GetAll() []interface{} {
	ret := make([]interface{}, q.size)
	switch {
	case q.tail > q.head: // 在一整个环内
		copy(ret, q.ele[q.head:])
	case q.tail < q.head: // 跨环
		copy(ret, q.ele[q.head:])
		copy(ret[q.size-q.head:], q.ele[:q.head])
	}
	return ret
}

func (q *CQueue) Clear() {
	for i := range q.ele {
		q.ele[i] = nil
	}
	q.head = 0
	q.tail = 0
}
