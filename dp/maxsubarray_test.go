package dp

import "testing"

func TestMaxSubArray(t *testing.T) {
	// 4, -1, 2, 1
	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	println(maxSubArray1(array))
	println(maxSubArray2(array))
}
