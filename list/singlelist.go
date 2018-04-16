package list

import (
	"fmt"
)

type SNode struct {
	data interface{} // 存储的数据
	next *SNode
}

type SingleList struct {
	size int
	head *SNode
	tail *SNode
}

func NewSingleList() *SingleList {
	return &SingleList{0, nil, nil}
}

func (l SingleList) IsEmpty() bool {
	return l.size == 0
	//return l.head.next == nil
}

func (l SingleList) Length() int {
	return l.size
}

func (l *SingleList) Add(data interface{}) bool {
	node := new(SNode)
	node.data = data
	if node == nil {
		return false
	}
	if l.size == 0 {
		l.head = node
	} else {
		tmp := l.tail
		tmp.next = node // 将新元素append到原来tail的后面
	}
	l.tail = node //更新list的tail
	l.size++
	return true
}

// 从1开始
func (l *SingleList) Insert(index int, data interface{}) bool {
	if data == nil || index > l.size || index < 1 { // index >=1   size>=0
		return false
	}
	node := &SNode{data: data,next:nil}
	if index == 1 {
		// AddFirst()
		node.next = l.head
		l.head = node
	} else {
		tmp := l.head
		for i := 1; i < index; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [new] next
		node.next = tmp.next
		tmp.next = tmp
	}
	l.size++
	return true
}

func (l SingleList) GetByIndex(index int) *SNode {
	if index > l.size {
		return nil
	}
	tmp := l.head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	return tmp
}

// 从1开始
func (l *SingleList) RemoveByIndex(index int) bool {
	if index > l.size || index < 1 || l.IsEmpty() {
		return false
	}
	if index == 1 {
		l.head = l.head.next
		if l.size == 1 {
			l.tail = nil
		}
	} else {
		tmp := l.head
		for i := 1; i < index; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [index] next
		tmp.next = tmp.next.next
		if index == l.size {
			l.tail = tmp
		}
	}
	l.size--
	return true
}

func (l *SingleList) RemoveData(data interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	node := l.head
	for node.next != nil {
		if node.data == data {
			node.next = node.next.next
			l.size--
			return true
		}
	}
	return false
}

func (l *SingleList) RemoveLast() {

}

func (l *SingleList) GetAll() []interface{} {
	obj := make([]interface{}, 0)
	if l.IsEmpty() {
		fmt.Println("list is empty")
		return nil
	}
	node := l.head
	for node.next != nil {
		obj = append(obj, node.data)
		node=node.next
	}
	obj = append(obj, node.data)
	return obj
}
