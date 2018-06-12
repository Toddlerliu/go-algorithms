package graph

// 邻接表（稀疏图，边少）
type AdjacencyList struct {
	vertexs  int  // 点数
	edges    int  // 边数
	directed bool // 是否有向
	//graph    [][]int // 和顶点相连的所有顶点编号
	graph map[int][]int // 和顶点相连的所有顶点编号
}

func NewSparseGraph(vers int, directed bool) *AdjacencyList {
	g := new(AdjacencyList)
	g.vertexs = vers
	g.edges = 0
	g.directed = directed
	g.graph = make(map[int][]int)
	for i := 0; i < vers; i++ {
		tmp := make([]int, 0)
		g.graph[i] = tmp
	}
	//g.graph = make([][]int, vers)
	//for i := 0; i < vers; i++ {
	//	tmp := make([]int, vers)
	//	g.graph[i] = tmp
	//}
	return g
}

func (g AdjacencyList) VersNum() int {
	return g.vertexs
}

func (g AdjacencyList) EdgeNum() int {
	return g.edges
}

// O(n)（二维数组）
func (g *AdjacencyList) AddEdge(v1, v2 int) {
	//if isConnected, err := g.hasEdge(v1, v2); err == nil {
	//	if isConnected {
	//		// 忽略平行边
	//		return
	//	}
	//	g.graph[v1] = append(g.graph[v1], v2)
	//	if v1 != v2 && !g.directed { // 处理自环边,且是无向图
	//		g.graph[v2] = append(g.graph[v2], v1)
	//	}
	//	g.edges++
	//}
	isConnected := g.hasEdge(v1, v2)
	if isConnected {
		// 忽略平行边
		return
	}
	g.graph[v1] = append(g.graph[v1], v2)
	if v1 != v2 && !g.directed { // 处理自环边,且是无向图
		g.graph[v2] = append(g.graph[v2], v1)
	}
	g.edges++
}

// 判断两顶点是否有边（连接）
// error为空参数非法；true连接；false不连接
// O(n)(二维数组)
func (g *AdjacencyList) hasEdge(v1, v2 int) bool {
	//if (v1 >= 0 && v1 < g.vertexs) && (v2 >= 0 && v2 < g.vertexs) {
	//	// O(n)
	//	for i := 0; i < len(g.graph[v1]); i++ {
	//		if g.graph[v1][i] == v2 {
	//			return true, nil
	//		}
	//	}
	//}
	if (v1 >= 0 && v1 < g.vertexs) && (v2 >= 0 && v2 < g.vertexs) {
		// O(m，m个)
		for _, v := range g.graph[v1] {
			if v == v2 {
				return true
			}
		}
	}
	return false
}

func (g AdjacencyList) AdjVertexs(v int) []int {
	return g.graph[v]
}
