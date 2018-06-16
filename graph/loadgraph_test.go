package graph

import (
	"fmt"
	"testing"
)

func TestLoadGraphFromFile(t *testing.T) {
	g := LoadGraphFromFile("SparseGraph", "g2.txt")
	if g, ok := g.(*AdjacencyList); ok {
		fmt.Println("点数为：", g.VersNum(), "边数为：", g.EdgeNum())
		fmt.Println("和顶点0相邻的点是：", g.AdjVertexs(0))
		fmt.Println("和顶点1相邻的点是：", g.AdjVertexs(1))
		fmt.Println("和顶点3相邻的点是：", g.AdjVertexs(3))
	}

	g1 := LoadGraphFromFile("DenseGraph", "g2.txt")
	if g1, ok := g1.(*AdjacencyMatrix); ok {
		fmt.Println("点数为：", g1.VersNum(), "边数为：", g1.EdgeNum())
		PrintMatrix(g1.graph)
		fmt.Println("和顶点0相邻的点是：", g1.AdjVertexs(0))
	}

}
