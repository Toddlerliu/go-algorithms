package dp

import (
	"algorithm/utils"
)

// 最长公共子序列：https://leetcode-cn.com/problems/longest-common-subsequence/
// 最长公共子串(连续)

// dp
func longestCommonSubstring1(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return 0
	}
	arr1 := []byte(str1)
	arr2 := []byte(str2)
	// dp[i,j] 是 str1[i-1] 与 str2[j-1] 的最长公共子串长度
	// 若 num1[i-1] == num2[j-1]，则 dp[i,j] = dp[i-1,j-1] + 1
	// 若 num1[i-1] != num2[j-1]，则 dp[i,j] = 0
	// 空间复杂度O(m*n)
	dp := make([][]int, len(arr1)+1)
	for i := range dp {
		dp[i] = make([]int, len(arr2)+1)
	}
	max := 0
	// 时间复杂度O(m*n)
	for i := 1; i <= len(arr1); i++ {
		for j := 1; j <= len(arr2); j++ {
			if arr1[i-1] == arr2[j-1] {
				// 当前值只和左上角有关，可优化为2*n数组or一维数组 O(min{m,n})
				dp[i][j] = dp[i-1][j-1] + 1
				max = utils.Max(dp[i][j], max)
			}
		}
	}
	return max
}

// dp
// 空间复杂度O(num1.len * num2,len)
// 时间复杂度O(num1.len * num2,len)
func longestCommonSubsequence2(num1, num2 []int) int {
	if num1 == nil || len(num1) == 0 {
		return 0
	}
	if num2 == nil || len(num2) == 0 {
		return 0
	}

	// dp[i,j] 是 nums1前i个元素 与 nums2前j个元素 的最长公共子序列长度
	// 若 num1[i-1] == num2[j-1]，则 dp[i,j] = dp[i-1,j-1] + 1
	// 若 num1[i-1] != num2[j-1]，则 dp[i,j] = max{dp[i-1,j],dp[i,j-1]}
	dp := make([][]int, len(num1)+1)
	for i := range dp {
		dp[i] = make([]int, len(num2)+1)
	}
	/**
	        	num1[i] 序列1
			---------
			| a | b |  dp矩阵
	num1[j]	| c | y |
			---------
		1、num1[i-1] == num2[j-1], y=a+1
		2、num1[i-1] != num2[j-1], y=max{c,b}
		优化：当前y的值，只和a、b、c有关，只需2*n的数组即可
	*/
	for i := 1; i <= len(num1); i++ {
		for j := 1; j <= len(num2); j++ {
			if num1[i-1] == num2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(num1)][len(num2)]
}

// 递归
// 空间复杂度O(min{num1,num2})
// 时间复杂度O(2^min{num1,num2})
func longestCommonSubsequence1(num1, num2 []int) int {
	if num1 == nil || len(num1) == 0 {
		return 0
	}
	if num2 == nil || len(num2) == 0 {
		return 0
	}
	return _longestCommonSubsequence1(num1, len(num1), num2, len(num2))
}

func _longestCommonSubsequence1(num1 []int, i int, num2 []int, j int) int {
	if i == 0 || j == 0 {
		return 0
	}
	if num1[i-1] == num2[j-1] {
		return _longestCommonSubsequence1(num1, i-1, num2, j-1) + 1
	}
	return utils.Max(
		_longestCommonSubsequence1(num1, i-1, num2, j),
		_longestCommonSubsequence1(num1, i, num2, j-1),
	)
}
