package nqueens

import "testing"

func TestNewNQueens2(t *testing.T) {

	queens_four := NewNQueens2(4)
	queens_four.PlcaeQueens()
	println(queens_four.getWays())

	queens_eight := NewNQueens2(8)
	queens_eight.PlcaeQueens()
	println(queens_eight.getWays())

}
