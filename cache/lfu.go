package cache

import (
	"sync"
	"container/list"
)

// LFU（Least Frequently Used）
// paper ：http://dhruvbird.com/lfu.pdf
// head  1  2  5  9    freqNode list
//       x  z  b  c    dataNode
//       y  a

// or：
// 3 map：（k-v）、（k-count）、（count-keySet）。
// get：O(1)；put：O(1)：Space: O(n)；

type dataNode struct {
	key      string
	val      interface{}
	freqNode *list.Element
}

type freqNode struct {
	freq int
	item map[*dataNode]struct{}
}

// doubly-linked list + hashmap
type LFUCache struct {
	sync.Mutex
	capacity  int                  // 容量
	freqNodes *list.List           // 频率节点list，每个节点包括 频率数+数据Map（freqNode）
	data      map[string]*dataNode // 数据存储
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{
		capacity:  cap,
		data:      make(map[string]*dataNode),
		freqNodes: list.New(),
	}
}

func (c LFUCache) IsEmpty() bool {
	c.Lock()
	defer c.Unlock()
	return len(c.data) == 0
}

func (c LFUCache) IsFull() bool {
	return c.isFull()
}

func (c LFUCache) isFull() bool {
	c.Lock()
	defer c.Unlock()
	return len(c.data) >= c.capacity
}

func (c LFUCache) Size() int {
	c.Lock()
	defer c.Unlock()
	return len(c.data)
}

func (c *LFUCache) Put(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	// 1 exist ?
	if v, ok := c.data[key]; ok {
		// exists: update
		v.val = value
	} else {
		// not exist: add
		// 2 full ?
		if c.isFull() {
			// delete
			c.evict(1)
		}
		// add
		dn := &dataNode{
			key: key,
			val: value,
		}

		if e := c.freqNodes.Front(); e == nil {
			// empty
			head := &freqNode{
				freq: 1,
				item: make(map[*dataNode]struct{}),
			}
			head.item[dn] = struct{}{}
			c.freqNodes.PushFront(head) // 构建freq=1
		} else {
			if fn := e.Value.(*freqNode); fn.freq == 1 {
				fn.item[dn] = struct{}{}
				dn.freqNode = e
			} else {
				head := &freqNode{
					freq: 1,
					item: make(map[*dataNode]struct{}),
				}
				head.item[dn] = struct{}{}
				c.freqNodes.PushFront(head) // 构建freq=1
			}

		}
		c.data[key] = dn // 数据添加
	}
}

// 淘汰频率最低的count个元素
func (c *LFUCache) evict(count int) {
	for i := 0; i < count; i++ {
		// freqNode ele {freq,itemMap}
		if ele := c.freqNodes.Front(); ele != nil {
			// freqNode map[*dataNode]struct{}
			for dataNode, _ := range ele.Value.(*freqNode).item {
				if i < count {
					delete(c.data, dataNode.key)                 // 数据删除
					delete(ele.Value.(*freqNode).item, dataNode) // 同频率map删除dataNode
					if len(ele.Value.(*freqNode).item) == 0 {
						c.freqNodes.Remove(ele) // 删除频率节点
					}
					i++
				}
			}
		}
	}
}

func (c *LFUCache) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()
	if dn, ok := c.data[key]; ok {
		c.increment(dn)
		return dn.val
	}
	return nil
}

// update
func (c *LFUCache) increment(dn *dataNode) {
	fnEle := dn.freqNode
	fn := fnEle.Value.(*freqNode)
	nextFreq := fn.freq + 1
	delete(fn.item, dn) // 原频率列表删除

	nextFnEle := fnEle.Next() // 下一个freqNode
	if nextFnEle == nil {
		nextFnEle = c.freqNodes.InsertAfter(&freqNode{
			freq: nextFreq,
			item: make(map[*dataNode]struct{}),
		}, fnEle)
	}
	nextFnEle.Value.(*freqNode).item[dn] = struct{}{}
	dn.freqNode = nextFnEle
}

func (c *LFUCache) Remove(key string) {

}
