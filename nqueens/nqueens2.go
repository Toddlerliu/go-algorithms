package nqueens

type NQueens2 struct {
	// [col]=true 某列是否有皇后
	cols []bool
	// 对角线是否有皇后
	// ↖↘ 索引：row-col+n-1
	lTop2rBottom []bool
	// ↗↙ 索引：wor+col
	rTop2lBottom []bool
	ways         int
}

func NewNQueens2(n int) *NQueens2 {
	if n < 1 {
		return nil
	}
	return &NQueens2{
		cols:         make([]bool, n),
		lTop2rBottom: make([]bool, n<<1-1), // 单方向 2n-1 条斜线
		rTop2lBottom: make([]bool, n<<1-1),
	}
}

func (r *NQueens2) PlcaeQueens() {
	r.place(0)
}

// 放置第 row 行的皇后
func (r *NQueens2) place(row int) {
	if row == len(r.cols) {
		r.ways++
		return
	}
	for col := 0; col < len(r.cols); col++ {
		// 列判断
		if r.cols[col] {
			// col 列有皇后
			continue
		}
		// 对角线判断
		if r.lTop2rBottom[row-col+len(r.cols)-1] {
			continue
		}
		if r.rTop2lBottom[row+col] {
			continue
		}
		// 放置row行皇后，继续下一行
		r.cols[col] = true
		r.lTop2rBottom[row-col+len(r.cols)-1] = true
		r.rTop2lBottom[row+col] = true
		r.place(row + 1)
		// 回溯：当前col列放置皇后，第一行无法放置皇后，回溯到当前行继续col++寻找新位置
		// 重置
		r.cols[col] = false
		r.lTop2rBottom[row-col+len(r.cols)-1] = false
		r.rTop2lBottom[row+col] = false
	}
	// 本行所有列都不能放置皇后，回溯
}

func (r NQueens2) getWays() int {
	return r.ways
}
