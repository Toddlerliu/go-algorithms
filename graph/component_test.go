package graph

import (
	"fmt"
	"testing"
)

func TestNewComponent(t *testing.T) {
	g := LoadGraphFromFile("SparseGraph", "g1.txt")
	c := NewComponent(g.(I))
	fmt.Println("g1的联通分量为：", c.Count())
	fmt.Println("0和1是否相连：", c.IsConnected(0, 1))
	fmt.Println("0和4是否相连：", c.IsConnected(0, 12))

	g1 := LoadGraphFromFile("SparseGraph", "g2.txt")
	c1 := NewComponent(g1.(I))
	fmt.Println("g2的联通分量为：", c1.Count())
}
