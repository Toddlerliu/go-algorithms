package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSparseGraph(t *testing.T) {
	fmt.Println("===sparse graph")
	g := NewDenseGraph(3, false)
	PrintMatrix(g.graph)
	fmt.Println(g.graph)
	fmt.Println(reflect.TypeOf(g.graph))

	//m := make(map[int]int{})
	m := map[int]int{}
	fmt.Println(reflect.TypeOf(m))

	//fmt.Println("vers:",g.VersNum(),"edges:",g.EdgeNum())
	//fmt.Println("添加边。。。")
	//g.AddEdge(0,1)
	//PrintMatrix(g.graph)
	//fmt.Println("vers:",g.VersNum(),"edges:",g.EdgeNum())
}
