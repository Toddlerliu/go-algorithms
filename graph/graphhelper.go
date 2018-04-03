package graph

import "fmt"

func PrintMatrix(m [][]int) {
	fmt.Println("矩阵如下：")
	for _, v := range m {
		fmt.Println(v)
	}
}
