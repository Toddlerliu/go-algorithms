package graph

import (
	"fmt"
	"testing"
)

func TestDenseGraph(t *testing.T) {
	g := NewDenseGraph(3, false)
	PrintMatrix(g.graph)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	fmt.Println("添加边。。。")
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	PrintMatrix(g.graph)
	fmt.Println("vers:", g.VersNum(), "edges:", g.EdgeNum())
	fmt.Println("和顶点0相邻的点是：", g.AdjVertexs(0))
}
