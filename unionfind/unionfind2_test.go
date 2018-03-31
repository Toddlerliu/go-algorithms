package unionfind

import (
	"fmt"
	"testing"
)

func TestNewUnionFind2(t *testing.T) {
	uf := NewUnionFind2(5)
	fmt.Println("uf:", uf)
	fmt.Println("1的parent:", uf.Find(1))
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("0和3连接中...")
	uf.Union(0, 3)
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
}
