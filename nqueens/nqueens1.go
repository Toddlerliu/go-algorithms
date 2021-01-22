package nqueens

import "math"

type NQueens struct {
	// [row]=col 皇后位置
	cols []int
	ways int
}

func NewNQueens(n int) *NQueens {
	if n < 1 {
		return nil
	}
	return &NQueens{
		cols: make([]int, n),
	}
}

func (r *NQueens) PlcaeQueens() {
	r.place(0)
}

// 放置第 row 行的皇后
func (r *NQueens) place(row int) {
	if row == len(r.cols) {
		r.ways++
		return
	}
	for col := 0; col < len(r.cols); col++ {
		if r.isValid(row, col) {
			// 放置row行皇后，继续下一行
			r.cols[row] = col
			r.place(row + 1)
			// 回溯：当前皇后位于row行col列，row+1行无法放置皇后，回溯到row行继续col++寻找新位置
		} // 				 	↑
	} // 					 	↑
	// 本行所有列都不能放置皇后，回溯 ↑
}

func (r *NQueens) isValid(row, col int) bool {
	for i := 0; i < row; i++ {
		// i行col列已有皇后
		if r.cols[i] == col {
			return false
		}
		// 斜线 → 斜率 (y1-y2)/(x1-x2)= (+/-)1
		if math.Abs(float64(col-r.cols[i])) == math.Abs(float64(row-i)) {
			return false
		}
	}
	return true
}

func (r NQueens) getWays() int {
	return r.ways
}
