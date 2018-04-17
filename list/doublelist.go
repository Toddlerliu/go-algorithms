package list

import "fmt"

type DNode struct {
	data interface{} // 存储的数据
	prev *DNode
	next *DNode
}

type DoublyLinkedList struct {
	size int
	head *DNode
	tail *DNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{0, nil, nil}
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return false
}

func (l *DoublyLinkedList) Size(data interface{}) int {
	return 0
}

func (l *DoublyLinkedList) Contains(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) Add(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) Insert(index int, data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) AddFirst(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) AddLast(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) AddAll(data []interface{}) bool {
	return false
}

func (l *DoublyLinkedList) InsertAll(index int, data []interface{}) bool {
	return false
}

// Replace
func (l *DoublyLinkedList) Set(index int, data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) GetFirst() *SNode {
	return nil
}

func (l *DoublyLinkedList) GetByIndex() *SNode {
	return nil
}

func (l *DoublyLinkedList) GetLast() *SNode {
	return nil
}

func (l *DoublyLinkedList) GetAll() []interface{} {
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

func (l *DoublyLinkedList) IndexOf(data interface{}) int {
	return 0
}

func (l *DoublyLinkedList) RemoveData(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) RemoveFirst() *SNode {
	return nil
}

func (l *DoublyLinkedList) RemoveLast() *SNode {
	return nil
}

func (l *DoublyLinkedList) RemoveByIndex(index int) *SNode {
	return nil
}

func (l *DoublyLinkedList) RemoveAll(data []interface{}) bool {
	return false
}

func (l *DoublyLinkedList) Clear() {

}
