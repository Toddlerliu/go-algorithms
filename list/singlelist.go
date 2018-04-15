package list

import "fmt"

type SNode struct {
	data interface{} // 存储的数据
	next *SNode
}

type SingleList struct {
	size int
	head *SNode
	tail *SNode
}

func NewSNode(data interface{}, next *SNode) *SNode {
	return &SNode{data, next}
}

func NewSingleList() *SingleList {
	return &SingleList{0, nil, nil}
}

func (l SingleList) IsEmpty() bool {
	//return l.size==0
	return l.head.next == nil
}

func (l SingleList) Length() int {
	return l.size
}

func (l *SingleList) Add(data interface{}) bool {
	node := new(SNode)
	node.data=data
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

func (l *SingleList) Insert(position int, node *SNode) bool {
	if node == nil || position > l.size || position < 1 { // posi >=1   size>=0
		return false
	}
	if position == 1 {
		// AddFirst()
		node.next = l.head
		l.head = node
	} else {
		tmp := l.head
		for i := 1; i < position; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [new] next
		node.next=tmp.next
		tmp.next=tmp
	}
	l.size++
	return true
}

func (l *SingleList) PrintSingleList() {
	if l.IsEmpty() {
		fmt.Println("list is empty")
		return
	}
	head := l.head
}