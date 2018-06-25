package graph

import (
	"fmt"
	"testing"
)

func TestNewShortestPath(t *testing.T) {
	g := LoadGraphFromFile("SparseGraph", "g2.txt")
	graph := NewShortestPath(g.(I), 0)
	fmt.Println("0到4的距离是：", graph.Length(4))
	fmt.Println("0到6的最短路径是：", graph.ShowPath(6))
}
