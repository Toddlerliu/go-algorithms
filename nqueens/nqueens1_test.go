package nqueens

import "testing"

func TestNewNQueens(t *testing.T) {

	queens_four := NewNQueens(4)
	queens_four.PlcaeQueens()
	println(queens_four.getWays())

	queens_eight := NewNQueens(8)
	queens_eight.PlcaeQueens()
	println(queens_eight.getWays())

}