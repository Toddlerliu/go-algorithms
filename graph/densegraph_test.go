package graph

import (
	"testing"
)

func TestDenseGraph(t *testing.T) {
	g := NewDenseGraph(3, false)
	g.graph[0][0] = 0
	g.graph[0][1] = 1
	g.graph[0][2] = 1
	g.graph[1][0] = 1
	g.graph[1][1] = 0
	g.graph[1][2] = 0
	g.graph[2][0] = 1
	g.graph[2][1] = 0
	g.graph[2][2] = 1

	PrintMatrix(g.graph)
}
