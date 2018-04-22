package list

type DNode struct {
	data interface{}
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
	return l.size == 0
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

// 若存在返回index
func (l *DoublyLinkedList) Contains(data interface{}) (int, bool) {
	index := 1
	for x := l.head; x.data != nil; x = x.next {
		if x.data == data {
			return index, true
		}
		index++
	}
	return -1, false
}

func (l *DoublyLinkedList) IndexOf(data interface{}) int {
	if i, ok := l.Contains(data); ok {
		return i
	}
	return -1
}

// Append、AddLast
func (l *DoublyLinkedList) Add(data interface{}) {
	node := &DNode{data: data}
	if l.size == 0 {
		l.head = node
		l.tail = node
		node.prev = nil
		node.next = nil
	} else {
		node.prev = l.tail
		node.next = nil
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *DoublyLinkedList) AddFirst(data interface{}) bool {
	return l.Insert(1, data)
}

// 从1开始
func (l *DoublyLinkedList) Insert(index int, data interface{}) bool {
	if data == nil || l.size < 1 || index < 1 || index > l.size+1 {
		return false
	}
	node := &DNode{data: data}
	if index == 1 {
		node.prev = nil
		node.next = l.tail
		l.tail.prev = node
		l.tail = node
	} else {
		tmp := l.head
		for i := 2; i < index; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [new] next
		node.prev = tmp
		node.next = tmp.next
		tmp.next.prev = node
		tmp.next = node
	}
	l.size++
	return true
}

func (l *DoublyLinkedList) AddAll(datas []interface{}) {
	len := len(datas)
	for i := 0; i < len; i++ {
		l.Add(datas[i])
	}
}

func (l *DoublyLinkedList) InsertAll(index int, datas []interface{}) bool {
	if datas == nil || l.size < 1 || index < 1 || index > l.size+1 {
		return false
	}
	// prev(tmp) [new] next
	if index == 1 {
		tmp := l.head
		node := &DNode{data: datas[0]}
		l.head = node
		for i := 2; i <= len(datas); i++ {
			l.Add(datas[i-1])


		}

	} else {
		tmp := l.head
		for i := 2; i < index; i++ {
			tmp = tmp.next
		}
		// prev(tmp) [new] next
		node.prev = tmp
		node.next = tmp.next
		tmp.next.prev = node
		tmp.next = node
	}
	l.size++
	return true
}

// Replace
func (l *DoublyLinkedList) Set(index int, data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) GetFirst() *DNode {
	return l.head
}

func (l *DoublyLinkedList) GetLast() *DNode {
	return l.tail
}

// 从1开始
func (l *DoublyLinkedList) GetByIndex() *DNode {
	return nil
}

func (l *DoublyLinkedList) GetAll() []interface{} {
	obj := make([]interface{}, 0)
	if l.IsEmpty() {
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

func (l *DoublyLinkedList) RemoveData(data interface{}) bool {
	return false
}

func (l *DoublyLinkedList) RemoveFirst() *DNode {
	tmp := l.head
	l.head=l.head.next
	l.size--
	return tmp
}

func (l *DoublyLinkedList) RemoveLast() *DNode {
	return nil
}

func (l *DoublyLinkedList) RemoveByIndex(index int) *DNode {
	return nil
}

func (l *DoublyLinkedList) RemoveAll(data []interface{}) bool {
	return false
}

func (l *DoublyLinkedList) Clear() {
	for x := l.head; x != nil; x = x.next {
		x.data = nil
		x.prev = nil
		x.next = nil
	}
	l.head = nil
	l.tail = nil
	l.size = 0
}
