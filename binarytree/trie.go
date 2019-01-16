package mybinarysearch

import "container/list"

type TrieNode struct {
	char rune // 字符
	// or []*TrieNode
	children map[rune]*TrieNode
	data     interface{}
	deep     int // 深度
	leaf     bool
	freq     int
}

type Trie struct {
	root *TrieNode
	size int // 节点个数
}

func NewTrieNode(char rune, deep int) *TrieNode {
	return &TrieNode{
		char:     char,
		children: make(map[rune]*TrieNode),
		deep:     deep,
		//freq:     1,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(' ', 1),
		size: 1,
	}
}

func (t *Trie) Insert(key string, data interface{}) {
	if len(key) == 0 {
		return
	}
	node := t.root
	chars := []rune(key)
	for _, char := range chars {
		kid, ok := node.children[char]
		if !ok {
			kid = NewTrieNode(char, node.deep+1)
			node.children[char] = kid
		}
		node.freq++
		node = kid
	}
	node.data = data
	node.leaf = true
}

// 前缀匹配
func (t *Trie) PrefixSearch(key string, limit int) (nodes []*TrieNode) {
	if len(key) == 0 {
		return
	}
	node := t.root
	chars := []rune(key)
	for _, char := range chars {
		kid, ok := node.children[char]
		if !ok {
			return
		}
		node = kid
	}

	// BFS
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		front := queue.Front()
		if node, ok := front.Value.(*TrieNode); ok {
			queue.Remove(front)
			if node.leaf == true {
				nodes = append(nodes, node)
				if len(nodes) > limit {
					return
				}
				continue
			}
			for _, kid := range node.children {
				queue.PushBack(kid)
			}
		}
	}
	return
}

// 全匹配
func (t *Trie) Search(key string) (bool, interface{}) {
	if len(key) == 0 {
		return false, nil
	}
	node := t.root
	chars := []rune(key)
	for _, char := range chars {
		kid, ok := node.children[char]
		if !ok {
			return false, nil
		}
		node = kid
	}
	if node.leaf == true {
		return true, node.data
	} else {
		return false, nil
	}
}

func (t *Trie) Remove(key string) {
	if len(key) == 0 {
		return
	}
	node := t.root
	chars := []rune(key)
	for _, char := range chars {
		if kid, ok := node.children[char]; ok {
			if kid.freq == 1 {
				delete(node.children, char)
				return
			}
			kid.freq--
			node = kid
		}
	}
	node.leaf = false
}
