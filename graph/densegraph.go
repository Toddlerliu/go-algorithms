package graph

// 邻接矩阵
type AdjacencyMatrix struct {
	vertexs  int     // 点数
	edges    int     // 边数
	directed bool    // 是否有向
	graph    [][]int // 矩阵(1:连接；0:不连接)
}

// 构造函数：参数：点数和是否有向
func NewDenseGraph(vers int, directed bool) *AdjacencyMatrix {
	g := new(AdjacencyMatrix)
	g.vertexs = vers
	g.edges = 0
	g.directed = directed
	g.graph = make([][]int, vers, vers)
	for i := 0; i < vers; i++ {
		tmp := make([]int, vers, vers)
		g.graph[i] = tmp
	}
	return g
}
