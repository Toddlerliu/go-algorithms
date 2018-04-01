package unionfind

import (
	"fmt"
	"testing"
)

func TestNewUnionFind5(t *testing.T) {
	fmt.Println("============UF5==========")
	uf := NewUnionFind5(5)
	fmt.Println("uf:", uf)
	fmt.Println("find(1):", uf.Find(1))
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("0和3连接中...")
	uf.Union(0, 3)
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("uf:", uf)
	fmt.Println("find2(1):", uf.Find2(1))
	fmt.Println("uf:", uf)
}
