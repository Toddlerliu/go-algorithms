package graph

import (
	"fmt"
	"testing"
)

func TestDenseGraph(t *testing.T) {
	g := NewDenseGraph(6, false)
	fmt.Println(g.graph)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	fmt.Println("添加边。。。")
	g.AddEdge(0, 1)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	PrintMatrix(g.graph)
	fmt.Println("和顶点0相邻的点是：", g.AdjVertexs(0))
}
