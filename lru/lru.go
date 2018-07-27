package lru

import (
	"container/list"
	"sync"
)

// 缓存淘汰算法：
// 思想：如果数据最近被访问过，那么将来被访问的几率也很高
// 当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰

//  实现方式：
//1.用一个数组来存储数据，给每一个数据项标记一个访问时间戳，每次插入新数据项的时候，先把数组中存在的数据项的时间戳自增，
//  并将新数据项的时间戳置为0并插入到数组中。每次访问数组中的数据项的时候，将被访问的数据项的时间戳置为0。当数组空间已满时，将时间戳最大的数据项淘汰。
//2.利用一个链表来实现，每次新插入数据的时候将新数据插到链表的头部；每次缓存命中（被访问），则将数据移到链表头部；
//  那么当链表满的时候，就将链表尾部的数据丢弃。
//3.双向链表和HashMap。插入数据项：如果新数据项在链表中存在（命中），则把该节点移到链表头部，如果不存在，
//  则新建一个节点，放到链表头部，若缓存满了，则把链表最后一个节点删除即可；访问数据：如果数据项在链表中存在，则把该节点移到链表头部。
//  HashMap读写时间复杂度O(1); 双向链表插入、删除时间复杂度O(1)。

type entry struct {
	key   interface{}
	value interface{}
}

type LRUCache struct {
	mu          sync.Mutex
	maxCapacity int // 最大存储数量
	//currCapacity int        // 当前存储数量
	list *list.List // entry→Element
	//type List struct {
	//	root Element // sentinel list element, only &root, root.prev, and root.next are used
	//	len  int
	//}
	cache map[interface{}]*list.Element // k-entry(Element)
	//type Element struct {
	//	next, prev *Element
	//	list *List //The list to which this element belongs
	//	Value interface{} // The value stored with this element
	//}
}

func NewLRUCache(maxCapacity int) *LRUCache {
	return &LRUCache{
		mu:          sync.Mutex{},
		maxCapacity: maxCapacity,
		//currCapacity: 0,
		list:  list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

func (l *LRUCache) Size() int {
	if l.cache == nil {
		return 0
	}
	//return l.currCapacity
	return l.list.Len()
}

// 根据key获取value（bool是否存在）
func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	if l.cache == nil {
		return nil, false
	}
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return nil, false
}

func (l *LRUCache) Set(key, value interface{}) {
	if l.cache == nil {
		l.list = list.New()
		l.cache = make(map[interface{}]*list.Element)
	}
	// 已存在,更新value
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		e.Value.(*entry).value = value
	}
	// 新增
	e := l.list.PushFront(&entry{key, value})
	l.cache[key] = e
	if l.list.Len() > l.maxCapacity {
		// 删除最后一个元素
		if l.cache == nil {
			l.RemoveOldest()
		}

	}
}

// 删除oldest元素
func (l *LRUCache) RemoveOldest() {
	if l.cache == nil {
		return
	}
	if e := l.list.Back(); e != nil {
		l.list.Remove(e)
		delete(l.cache, e.Value.(*entry).key)
	}
}

// 删除指定元素
func (l *LRUCache) Remove(key interface{}) {
	if l.cache == nil {
		return
	}
	if e, ok := l.cache[key]; ok {
		l.list.Remove(e)
		delete(l.cache, e.Value.(*entry).key)
	}
}
