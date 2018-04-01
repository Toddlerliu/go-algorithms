package unionfind

import (
	"fmt"
	"testing"
)

func TestNewUnionFind4(t *testing.T) {
	fmt.Println("============UF4==========")
	uf := NewUnionFind4(5)
	fmt.Println("uf:", uf)
	fmt.Println("1的parent:", uf.Find(1))
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("0和3连接中...")
	uf.Union(0, 3)
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("uf:", uf)
}
