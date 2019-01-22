package hash

import (
	"strconv"
	"sort"
	"errors"
	"hash/crc32"
)

const DefaultVirtualNodes = 20

type ConsistentHash struct {
	nodes         map[string]struct{} // 真实节点
	circle        map[uint32]string   // 哈希环(virtual)  hash(host)-host
	sortedNodes   []uint32            // 有序哈希环(virtual)，可用红黑树替代，eg：TreeMap
	replicaNumber int                 // 每个真实节点虚拟多少个虚拟节点
}

func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		nodes:         make(map[string]struct{}),
		circle:        make(map[uint32]string),
		sortedNodes:   make([]uint32, 0),
		replicaNumber: DefaultVirtualNodes,
	}
}

func (c *ConsistentHash) Add(hosts ...string) {
	for _, host := range hosts {
		if _, ok := c.nodes[host]; ok {
			return
		}
		// 填充真实节点
		c.nodes[host] = struct{}{}
		// 填充虚拟节点
		for i := 0; i < c.replicaNumber; i++ {
			virtualKey := virtualHash(i, host)
			c.circle[virtualKey] = host
			c.sortedNodes = append(c.sortedNodes, virtualKey)
		}
		// sort
		c.resort()
	}
}

func (c *ConsistentHash) resort() {
	// quickSort
	sort.Slice(c.sortedNodes, func(i, j int) bool {
		return c.sortedNodes[i] < c.sortedNodes[j]
	})
}

// 获取最近节点
func (c *ConsistentHash) Get(key string) (string, error) {
	if len(key) == 0 || len(c.circle) == 0 {
		return "", errors.New("empty")
	}
	hashKey := hash(key)
	// nearby node
	var node string
	if _, ok := c.circle[hashKey]; !ok {
		for _, v := range c.sortedNodes {
			// 顺时针第一个
			if v > hashKey {
				return c.circle[v], nil
			}
		}
		// 环上没有，环第一个
		return c.circle[c.sortedNodes[0]], nil
	}
	node = c.circle[hashKey]
	return node, nil
}

func (c *ConsistentHash) Remove(hosts ...string) {
	for _, host := range hosts {
		if _, ok := c.nodes[host]; ok {
			// 删除真实节点
			delete(c.nodes, host)
			for i := 0; i < c.replicaNumber; i++ {
				virtualKey := virtualHash(i, host)
				// 删除虚拟节点
				delete(c.circle, virtualKey)
				for i, val := range c.sortedNodes {
					if val == virtualKey {
						// 删除排序虚拟节点
						c.sortedNodes = append(c.sortedNodes[:i], c.sortedNodes[i+1:]...)
						continue
					}
				}
			}
		}
	}
}

func hash(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func virtualHash(idx int, key string) uint32 {
	key = key + strconv.Itoa(idx)
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}
