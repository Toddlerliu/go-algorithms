package graph

import (
	"fmt"
	"testing"
)

func TestNewPath(t *testing.T) {
	g := LoadGraphFromFile("SparseGraph", "g2.txt")
	graph := NewPath(g.(I), 0)
	fmt.Println("每个点是从哪里来的：", graph.from)
	fmt.Println("0到6的path是：", graph.ShowPath(6))
}
