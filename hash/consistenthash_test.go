package hash

import (
	"testing"
	"fmt"
)

func TestConsistentHash(t *testing.T) {
	nodes := []string{"192.168.0.1:8080", "192.168.0.2:8080", "192.168.0.3:8080", "192.168.0.4:8080", "192.168.0.4:8080"}
	hash := NewConsistentHash()
	hash.Add(nodes...)

	node, _ := hash.Get("aaa")
	fmt.Printf("%s的所在节点为：%s\n", "aaa", node)
	node, _ = hash.Get("bbb")
	fmt.Printf("%s的所在节点为：%s\n", "bbb", node)
	node, _ = hash.Get("ccc")
	fmt.Printf("%s的所在节点为：%s\n", "ccc", node)
	node, _ = hash.Get("ddd")
	fmt.Printf("%s的所在节点为：%s\n", "ddd", node)
	fmt.Println("删除节点192.168.0.1:8080")
	hash.Remove("192.168.0.1:8080")
	node, _ = hash.Get("aaa")
	fmt.Printf("%s的所在节点为：%s\n", "aaa", node)
	node, _ = hash.Get("bbb")
	fmt.Printf("%s的所在节点为：%s\n", "bbb", node)
	node, _ = hash.Get("ccc")
	fmt.Printf("%s的所在节点为：%s\n", "ccc", node)
	node, _ = hash.Get("ddd")
	fmt.Printf("%s的所在节点为：%s\n", "ddd", node)

}
