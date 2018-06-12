package graph

import (
	"fmt"
	"testing"
)

func TestSparseGraph(t *testing.T) {
	fmt.Println("===sparse graph")
	g := NewSparseGraph(3, false)
	fmt.Println(g.graph)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	fmt.Println("添加边。。。")
	g.AddEdge(0, 1)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	fmt.Println(g.graph)
	fmt.Println("和顶点0相邻的点是：", g.AdjVertexs(0))
}
