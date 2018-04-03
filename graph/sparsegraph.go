package graph

// 邻接表
type AdjacencyList struct {
	vertexs  int  // 点数
	edges    int  // 边数
	directed bool // 是否有向
	graph    map[int][]int
	//graph    [][]int // 和顶点相连的所有顶点编号
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
	return g
}
