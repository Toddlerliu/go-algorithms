package maze

import (
	"testing"
	"fmt"
)

func TestMaze(t *testing.T) {
	maze := readMaze("maze.txt")

	start, end := point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	// 到终点走了多少步
	stepNum, _ := end.at(steps)
	fmt.Printf("从start到end，一共走了 %d 步", stepNum)
	fmt.Println()

	path, ok := path(steps, start, end)
	if ok {
		fmt.Println("path from start to end :", path)
	}
}
