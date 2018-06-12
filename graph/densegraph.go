package graph

import "errors"

// 邻接矩阵(稠密图，边多，eg：完全图（所有点都相连（边））)
type AdjacencyMatrix struct {
	vertexs  int     // 点数
	edges    int     // 边数
	directed bool    // 是否有向
	graph    [][]int // 矩阵(1:连接；0:不连接)
}

// 构造函数：参数：点数（0至n-1）和是否有向
func NewDenseGraph(vers int, directed bool) *AdjacencyMatrix {
	g := new(AdjacencyMatrix)
	g.vertexs = vers
	g.edges = 0
	g.directed = directed
	g.graph = make([][]int, vers)
	for i := 0; i < vers; i++ {
		tmp := make([]int, vers)
		g.graph[i] = tmp
	}
	return g
}

func (g AdjacencyMatrix) VersNum() int {
	return g.vertexs
}

func (g AdjacencyMatrix) EdgeNum() int {
	return g.edges
}

// 连接v1，v2两个顶点
func (g *AdjacencyMatrix) AddEdge(v1, v2 int) {
	if isConnected, err := g.hasEdge(v1, v2); err == nil {
		if isConnected {
			// 忽略平行边(邻接矩阵 无平行边)
			return
		}
		g.graph[v1][v2] = 1
		if !g.directed {
			// 无向图
			g.graph[v2][v1] = 1
		}
		g.edges++
	}
}

// 判断两顶点是否有边（连接）
// error为空参数非法；true连接；false不连接
// O(1)
func (g *AdjacencyMatrix) hasEdge(v1, v2 int) (bool, error) {
	if (v1 >= 0 && v1 < g.vertexs) && (v2 >= 0 && v2 < g.vertexs) {
		return g.graph[v1][v2] == 1, nil
	}
	return false, errors.New("error input")
}

func (g AdjacencyMatrix) AdjVertexs(v int) (slice []int) {
	if v >= 0 && v < g.vertexs {
		for i, isConnected := range g.graph[v] {
			if isConnected == 1 {
				slice = append(slice, i)
			}
		}
	}
	return slice
}
