package graph

// 图的连通分量

// 引用interface{}，充当泛型
type I interface {
	VersNum() int
	EdgeNum() int
	AddEdge(v1, v2 int)
	//hasEdge(v1, v2 int) bool
	AdjVertexs(v int) (slice []int)
}

type Component struct {
	graph   I      // 图
	visited []bool // 节点是否被访问过
	ccount  int    // 连通分量个数
	id      []int  // 两节点相连：id相同
}

func NewComponent(graph I) *Component {
	g := new(Component)
	vers := graph.VersNum()
	g.graph = graph
	g.visited = make([]bool, vers)
	g.id = make([]int, vers)
	for i, _ := range g.id {
		g.id[i] = -1
	}
	g.ccount = 0

	for i := 0; i < vers; i++ {
		if !g.visited[i] {
			// 未被遍历过
			g.dfs(i)
			g.ccount++
		}
	}
	return g
}

// 将和i相连接的节点遍历，没遍历的节点在另外的连接分量中
func (c *Component) dfs(i int) {
	c.visited[i] = true
	c.id[i] = c.ccount // 遍历相同的点，ccount相同 0 1 2 ..
	for _, v := range c.graph.AdjVertexs(i) {
		if !c.visited[v] {
			c.dfs(v)
		}
	}
}

func (c Component) Count() int {
	return c.ccount
}

func (c Component) IsConnected(m, n int) bool {
	if (m >= 0 && m < c.graph.VersNum()) && (n >= 0 && m < c.graph.VersNum()) {
		return c.id[m] == c.id[n]
	}
	return false
}
