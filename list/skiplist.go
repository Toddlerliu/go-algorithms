package list

import (
	"sync"
	"math/rand"
	"time"
)

// O(n/2)

const (
	DEFAULT_MAX_LEVEL = 8
)

type skiplistNode struct {
	key   int
	value interface{}
	//backward *skiplistNode   // 后退指针
	forward []*skiplistNode // 前进指针（跟DoublyLinkedList相比，额外的存储空间，空间换时间）；后继 or 后继后继后继...  跳跃
}

type SkipList struct {
	sync.RWMutex
	root   *skiplistNode // head节点
	level  int           // 层数
	Length int           // 当前长度（节点个数）
	rand   *rand.Rand
}

func newSkiplistNode(key int, value interface{}) *skiplistNode {
	return &skiplistNode{
		key:     key,
		value:   value,
		forward: make([]*skiplistNode, DEFAULT_MAX_LEVEL),
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		root: &skiplistNode{
			forward: make([]*skiplistNode, DEFAULT_MAX_LEVEL),
		},
		level: DEFAULT_MAX_LEVEL,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l *SkipList) isEmpty() bool {
	return l.Length == 0
}

func (l *SkipList) randLevel() int {
	return l.rand.Intn(DEFAULT_MAX_LEVEL)
}

func (l *SkipList) Insert(key int, value interface{}) bool {
	// 先查找到key的前一个位置
	node := l.root
	for i := l.level - 1; i >= 0; i-- {
		for node.forward != nil && node.forward[i].key < key {
			node = node.forward[i]
		}
	}

}

func (l *SkipList) Search(key int) (bool, interface{}) {
	l.RLock()
	defer l.RUnlock()

	node := l.root
	for i := l.level - 1; i >= 0; i-- { // ↓
		for node.forward != nil && node.forward[i].key < key { // 跳跃
			node = node.forward[i]
		}
	}
	node = node.forward[0]
	if node.key == key {
		return true, node.value
	}
	return false, nil
}
