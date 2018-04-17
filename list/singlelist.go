package list

import (
	"fmt"
)

type SNode struct {
	data interface{} // 存储的数据
	next *SNode
}

type SingleLinkedList struct {
	size int
	head *SNode
	tail *SNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{0, nil, nil}
}

func (l SingleLinkedList) IsEmpty() bool {
	return l.size == 0
	//return l.head.next == nil
}

func (l SingleLinkedList) Length() int {
	return l.size
}

func (l *SingleLinkedList) Add(data interface{}) bool {
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

func (l *SingleLinkedList) AddFirst(data interface{}) bool {
	return l.Insert(1, data)
}

func (l *SingleLinkedList) AddLast(data interface{}) bool {
	return l.Insert(l.size+1, data)
}

// 从1开始
func (l *SingleLinkedList) Insert(index int, data interface{}) bool {
	if data == nil || index > l.size+1 || index < 1 { // index >=1   size>=0
		return false
	}
	node := &SNode{data: data, next: nil}
	if index == 1 {
		// AddFirst()
		node.next = l.head
		l.head = node
	} else {
		tmp := l.head
		for i := 2; i < index; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [new] next
		node.next = tmp.next
		tmp.next = node
	}
	l.size++
	return true
}

func (l SingleLinkedList) GetByIndex(index int) *SNode {
	if index > l.size {
		return nil
	}
	tmp := l.head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	return tmp
}

func (l SingleLinkedList) GetFirst() *SNode {
	return l.head
}

func (l SingleLinkedList) GetLast() *SNode {
	return l.tail
}

func (l *SingleLinkedList) RemoveFirst() bool {
	return l.RemoveByIndex(1)
}

func (l *SingleLinkedList) RemoveLast() bool {
	return l.RemoveByIndex(l.size)
}

// 从1开始
func (l *SingleLinkedList) RemoveByIndex(index int) bool {
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
		for i := 2; i < index; i++ {
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

// 重复，删除第一个
func (l *SingleLinkedList) RemoveData(data interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	node := l.head
	if node.data != data {
		for node.next.data != data {
			node = node.next
		}
		node.next = node.next.next
		l.size--
		return true
	} else {
		// delete head
		l.head = l.head.next
		l.size--
		return true
	}
	return false
}

func (l *SingleLinkedList) GetAll() []interface{} {
	obj := make([]interface{}, 0)
	if l.IsEmpty() {
		fmt.Println("list is empty")
		return nil
	}
	node := l.head
	for node.next != nil {
		obj = append(obj, node.data)
		node = node.next
	}
	obj = append(obj, node.data)
	return obj
}
