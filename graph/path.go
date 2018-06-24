package graph

import (
	"algorithms/utils"
	"errors"
	"sort"
)

type path struct {
	graph   I      // 图
	source  int    // 起始点
	visited []bool // 节点i是否被访问过
	from    []int  // 记录从哪个节点过来 from[i]=v : i节点是从v节点访问过来
}

func NewPath(graph I, source int) *path {
	if source < 0 || source >= graph.VersNum() {
		return nil
	}
	p := new(path)
	p.source = source
	p.graph = graph
	p.visited = make([]bool, graph.VersNum())
	p.from = make([]int, graph.VersNum())
	for i, _ := range p.from {
		p.from[i] = -1
	}
	p.dfs(source)
	return p
}

func (p *path) dfs(i int) {
	p.visited[i] = true
	sort.Ints(p.graph.AdjVertexs(i))
	for _, v := range p.graph.AdjVertexs(i) { // 与i相邻的节点
		if !p.visited[v] {
			p.from[v] = i // from i to v， v节点是从i节点访问过来
			p.dfs(v)
		}
	}
}

func (p path) HasPath(dst int) (bool, error) {
	if dst >= 0 && dst < p.graph.VersNum() {
		return p.visited[dst], nil
	}
	return false, errors.New("wrong index")
}

func (p path) ShowPath(dst int) []int {
	return p.path(dst)
}

// path from source to dst
func (p *path) path(dst int) (res []int) {
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
