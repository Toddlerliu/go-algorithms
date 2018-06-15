package list

type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

// List
// Queue
// Stack
type DoublyLinkedList struct {
	size int
	head *DNode
	tail *DNode
}

// 可重复
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
	if l.size > 0 {
		for x := l.head; x != nil; x = x.next {
			if x.data == data {
				return index, true
			}
			index++
		}
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
		node.next = l.head
		l.head.prev = node
		l.head = node
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
	i := index
	for _, v := range datas {
		l.Insert(i, v)
		i++
	}
	return true
}

// Replace
func (l *DoublyLinkedList) Set(index int, data interface{}) bool {
	if data == nil || l.size < 1 || index < 1 || index > l.size+1 {
		return false
	}
	tmp := l.head
	for i := 1; i <= index; i++ {
		if index == 1 {
			tmp.data = data
			return true
		}
		tmp = tmp.next
	}
	tmp.data = data
	return true
}

func (l *DoublyLinkedList) GetFirst() *DNode {
	return l.head
}

// Stack 出栈不删除(Peek)
func (l *DoublyLinkedList) GetLast() *DNode {
	return l.tail
}

// 从1开始
func (l *DoublyLinkedList) GetByIndex(index int) *DNode {
	if l.size < 1 || index < 1 || index > l.size+1 {
		return nil
	}
	tmp := l.head
	for i := 1; i < index; i++ {
		if index == 1 {
			return tmp
		}
		tmp = tmp.next
	}
	return tmp
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
	for x := l.head; x != nil; x = x.next {
		if x.data == data {
			// first
			if x.prev == nil {
				l.RemoveFirst()
				return true
			}
			// last
			if x.next == nil {
				//x.prev.next = nil
				//l.tail = x.prev
				//l.size--
				l.RemoveLast()
				return true
			}
			x.prev.next = x.next
			x.next.prev = x.prev
			l.size--
			return true
		}
	}
	return false

	//x := l.head
	//for i := 1; i <= l.size; i++ {
	//	if x.data == data {
	//		// first
	//		if x.prev == nil {
	//			l.RemoveFirst()
	//			return true
	//		}
	//		// last
	//		if x.next == nil {
	//			l.RemoveLast()
	//			return true
	//		}
	//		x.prev.next = x.next
	//		x.next.prev = x.prev
	//		l.size--
	//		return true
	//	}
	//	x = x.next
	//}
	//return false
}

func (l *DoublyLinkedList) RemoveFirst() *DNode {
	if l.size > 0 {
		tmp := l.head
		l.head.next.prev = nil
		l.head = l.head.next
		l.size--
		return tmp
	}
	return nil
}

func (l *DoublyLinkedList) RemoveLast() *DNode {
	tmp := l.tail
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	l.size--
	return tmp
}

func (l *DoublyLinkedList) RemoveByIndex(index int) *DNode {
	if l.size < 1 || index < 1 || index > l.size+1 {
		return nil
	}
	tmp := l.head
	for i := 1; i <= index; i++ {
		if index == 1 {
			return l.RemoveFirst()
		}
		if tmp.next == nil {
			return l.RemoveLast()
		}
		if i == index {
			tmp.prev.next = tmp.next
			tmp.next.prev = tmp.prev
			l.size--
		}
		tmp = tmp.next
	}
	return tmp
}

func (l *DoublyLinkedList) RemoveAll(data []interface{}) {
	for _, v := range data {
		l.RemoveData(v)
	}
}

// Queue 出队不删除
func (l DoublyLinkedList) QueuePeek() *DNode {
	return l.GetFirst()
}

// Queue 出队并删除
func (l *DoublyLinkedList) Poll() *DNode {
	return l.RemoveFirst()
}

// Stack 出栈不删除
func (l DoublyLinkedList) StackPeek() *DNode {
	return l.GetLast()
}

// Stack 出栈并删除
func (l *DoublyLinkedList) Pop() *DNode {
	return l.RemoveLast()
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
