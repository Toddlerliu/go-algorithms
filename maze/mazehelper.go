package maze

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	//filepath := "src\\algorithms\\maze\\" + filename
	filepath := "G:\\Code\\goAlgorithms\\src\\algorithms\\maze\\" + filename
	fmt.Println(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func ReverseInt(s []point) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}
