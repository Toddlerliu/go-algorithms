package dp

import "algorithm/utils"

/**
最长递增子序列(非连续):https://leetcode-cn.com/problems/longest-increasing-subsequence/
*/

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// dp[i]以nums[i]结尾的最长上升子序列的长度
	// dp[i] = max(dp[j<i](num[j]<num[i])+1, dp[i])
	// 空间复杂度：O(n)
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	// 时间复杂度：O(n^2)
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			dp[i] = utils.Max(dp[j]+1, dp[i])
		}
		max = utils.Max(dp[i], max)
	}
	return max
}
