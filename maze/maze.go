package maze

type point struct {
	i, j int // i：↓；j：→
}

var dirs = [4]point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 点在二维数组（maze，steps）中的值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

//   |  0  1  2  3  4
//---|----------------
// 0 |  0     4  5  6
// 1 |  1  2  3     7
// 2 |  2     4     8
// 3 |           10 9
// 4 |        12 11
// 5 |        13 12 13
//
func walk(maze [][]int, start, end point) [][]int {
	// 路径：元素代表 从start走了多少步才走到这一格
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:] // 出队

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze中0的才能走，1是墙
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// step为0才能走，不为0表示走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 不能回原点
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func path(maze [][]int, start, end point) ([]point, bool) {

	if start.i < 0 || start.i >= len(maze) || start.j < 0 || start.j >= len(maze[start.i]) {
		return nil, false
	}
	if end.i < 0 || end.i >= len(maze) || end.j < 0 || end.j > len(maze[end.i]) {
		return nil, false
	}

	endValue := maze[end.i][end.j]
	cur := end

	path := make([]point, 0)
	path = append(path, end)
	tmp := endValue

	if endValue > 0 {
		// 通路
		for i := tmp - 1; i >= 0; i-- {
			tmp -= 1
			for _, dir := range dirs {
				next := cur.add(dir) // 逆着推

				val, ok := next.at(maze)
				if !ok || val == 0 {
					continue
				}

				if maze[next.i][next.j] == tmp {
					path = append(path, next)
					cur = point{next.i, next.j}
				}
			}
		}
		path = append(path, start)
		ReverseInt(path)
		return path, true
	} else {
		// 非通路
		return nil, false
	}
	return nil, false
}
