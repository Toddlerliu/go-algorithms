package dp

import (
	"algorithm/utils"
)

/**
最大连续子序列和
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
*/
func maxSubArray1(array []int) int {
	if array == nil || len(array) == 0 {
		return 0
	}
	// 1、定义状态、状态转移方程
	// dp[i]以i结尾的最大连续子序列和
	// dp[i] = dp[i-1] + array[i]
	// 空间复杂度：O(n)
	dp := make([]int, len(array))
	// 2、初始状态
	dp[0] = array[0]
	// 3、最终解
	max := dp[0]
	// 时间复杂度：O(n)
	for i := 1; i < len(array); i++ {
		if dp[i-1] < 0 {
			dp[i] = array[i]
		} else {
			dp[i] = dp[i-1] + array[i]
		}
		max = utils.Max(dp[i], max)
	}
	return max
}

func maxSubArray2(array []int) int {
	if array == nil || len(array) == 0 {
		return 0
	}
	// 空间复杂度：O(1)
	dp := array[0]
	max := dp
	// 时间复杂度：O(n)
	for i := 1; i < len(array); i++ {
		if dp < 0 {
			dp = array[i]
		} else {
			dp = dp + array[i]
		}
		max = utils.Max(dp, max)
	}
	return max
}
