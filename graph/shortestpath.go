package graph

import (
	"algorithms/queue"
	"algorithms/utils"
	"errors"
	"sort"
)

// 无权图最短路径

type shortestpath struct {
	graph   I      // 图
	source  int    // 根节点
	visited []bool // 节点i是否被访问过
	from    []int  // 记录从哪个节点过来 from[i]=v : i节点是从v节点访问过来
	ord     []int  // 从source到每个节点的最短距离
}

func NewShortestPath(graph I, source int) *shortestpath {
	if source < 0 || source >= graph.VersNum() {
		return nil
	}
	spath := new(shortestpath)
	spath.source = source
	spath.graph = graph
	spath.visited = make([]bool, graph.VersNum())
	spath.from = make([]int, graph.VersNum())
	spath.ord = make([]int, graph.VersNum())
	for i, _ := range spath.from {
		spath.from[i] = -1
		spath.ord[i] = -1
	}

	// 广度优先遍历
	queue := queue.NewQueue()
	queue.Offer(source)
	spath.ord[source] = 0
	for !queue.Empty() {
		if i, ok := queue.Poll().(int); ok {
			sort.Ints(spath.graph.AdjVertexs(i))
			for _, v := range spath.graph.AdjVertexs(i) { // 与i相邻的节点
				if !spath.visited[v] {
					queue.Offer(v)
					spath.visited[v] = true
					spath.from[v] = i // from i to v， v节点是从i节点访问过来
					spath.ord[v] = spath.ord[i] + 1
				}
			}
		}
	}
	return spath
}

func (p shortestpath) HasPath(dst int) (bool, error) {
	if dst >= 0 && dst < p.graph.VersNum() {
		return p.visited[dst], nil
	}
	return false, errors.New("wrong index")
}

func (p shortestpath) ShowPath(dst int) []int {
	return p.path(dst)
}

// path from source to dst
func (p *shortestpath) path(dst int) (res []int) {
	// 倒推
	res = append(res, dst)
	for p.from[dst] != -1 { // -1 是source
		res = append(res, p.from[dst])
		dst = p.from[dst]
	}
	// 反转
	utils.ReverseInt(res)
	return
}

func (p *shortestpath) Length(dst int) int {
	if dst >= 0 && dst < p.graph.VersNum() {
		return p.ord[dst]
	}
	return -1
}
