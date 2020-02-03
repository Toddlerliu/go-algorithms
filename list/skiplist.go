package list

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DEFAULT_MAX_LEVEL = 8
	INT_MAX           = int(^uint(0) >> 1)
	INT_MIN           = ^INT_MAX
)

type skiplistNode struct {
	key                   int
	value                 interface{}
	up, down, left, right *skiplistNode
}

func newSkiplistNode(key int, value interface{}) *skiplistNode {
	return &skiplistNode{
		key:   key,
		value: value,
	}
}

type SkipList struct {
	head   *skiplistNode
	tail   *skiplistNode
	height int // 索引高度，0为数据
	size   int // 当前长度（节点个数）
	rand   *rand.Rand
}

func NewSkipList() *SkipList {
	head := &skiplistNode{
		key:   INT_MIN,
		value: nil,
	}
	tail := &skiplistNode{
		key:   INT_MAX,
		value: nil,
	}
	head.right = tail
	tail.left = head
	return &SkipList{
		head:   head,
		tail:   tail,
		height: 0,
		size:   0,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l *SkipList) isEmpty() bool {
	return l.size == 0
}

// 随机节点层数
// 1、level=0：数据链表
// 2、如果一个节点有第i层(i>0)指针（即节点已经在第0层到第i-1层链表中），那么它有第(i+1)层指针的概率为p。
// 3、level<=MaxLevel-1：节点最大的层数最大值
// 平衡树每个节点包含2个指针（左/右子树）
// skiplist每个节点包含的指针数目平均为1/(1-p)，p=1/4（Redis的选择）
// 1/(1-p) :
func (l *SkipList) randLevel() int {
	level := 0
	p := 0.25
	for l.rand.Float64() < p && level < DEFAULT_MAX_LEVEL {
		level++
	}
	return level
}

// O(logN)
func (l *SkipList) Search(key int) *skiplistNode {
	node := l.head
	for {
		// 右移
		for node.right.key != INT_MAX && node.right.key <= key {
			node = node.right
		}
		// 下移
		if node.down != nil {
			node = node.down
		} else {
			break
		}
	}
	// node <= key < node.right
	return node
}

func (l *SkipList) Contains(key int) bool {
	node := l.Search(key)
	return node.key == key
}

func (l *SkipList) Get(key int) interface{} {
	node := l.Search(key)
	if node.key == key {
		return node.value
	}
	return nil
}

func (l *SkipList) get(key int) *skiplistNode {
	node := l.Search(key)
	if node.key == key {
		return node
	}
	return nil
}

// 新节点和各层索引节点逐一比较，确定原链表的插入位置。O(logN)
// 把索引插入到原链表。O(1)
// 利用抛硬币的随机方式，决定新节点是否提升为上一级索引。O(logN)
// 插入时间复杂度: O(logN)
// 空间复杂度：O(N) (n/2 + n/4 + n/8 + … + 8 + 4 + 2 = n-2)
func (l *SkipList) Insert(key int, value interface{}) {
	// 先查找到key的前一个位置
	nearNode := l.Search(key)
	newNode := newSkiplistNode(key, value)
	newNode.left = nearNode
	newNode.right = nearNode.right
	nearNode.right.left = newNode
	nearNode.right = newNode

	// 构建k级索引,从1开始
	level := l.randLevel()
	fmt.Println("level is :", level)
	currentLevel := 0
	if level > 0 {
		for i := 1; i <= level; i++ {
			if currentLevel >= l.height {
				// 加层
				l.height++
				upperHead := newSkiplistNode(INT_MIN, nil)
				upperTail := newSkiplistNode(INT_MAX, nil)

				upperHead.right = upperTail
				upperHead.down = l.head
				l.head.up = upperHead

				upperTail.left = upperHead
				upperTail.down = l.tail
				l.tail.up = upperTail

				l.head = upperHead
				l.tail = upperTail
			}

			for nearNode != nil && nearNode.up == nil {
				nearNode = nearNode.left
			}
			nearNode = nearNode.up
			upNode := newSkiplistNode(key, value)
			upNode.left = nearNode
			upNode.right = nearNode.right
			upNode.down = newNode
			nearNode.right.left = upNode
			nearNode.right = upNode
			newNode.up = upNode
			newNode = upNode

			// 构建上层索引
			currentLevel++
		}
	}
	l.size++
}

// O(logN)
// 自上而下，查找第一次出现节点的索引，并逐层找到每一层对应的节点。O(logN)
// 删除每一层查找到的节点，如果该层只剩下1个节点，删除整个一层（原链表除外）。O(logN)
// 删除时间复杂度是O(logN)。
func (l *SkipList) delete(key int) {
	// 0 层数据
	node := l.get(key)
	for node != nil {
		left := node.left
		right := node.right
		left.right = right
		right.left = left

		node = node.up
	}
}

// 层级打印
func (l *SkipList) levelPrint() {
	tmp := l.head
	i := l.height
	for tmp != nil {
		if i != 0 {
			fmt.Printf("第%d层的索引为：", i)
		} else {
			fmt.Printf("第%d层的数据为：", i)
		}
		fmt.Println()
		node := tmp.right
		for node.right.key != INT_MAX {
			fmt.Printf(" -> %d", node.right.key)
			node = node.right
		}
		tmp = tmp.down
		i--
		fmt.Println()
	}
}
