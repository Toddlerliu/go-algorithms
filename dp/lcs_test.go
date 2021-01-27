package dp

import "testing"

func TestLCS(t *testing.T) {
	num1 := []int{1, 3, 5, 9, 10}
	num2 := []int{1, 3, 9, 10}
	println(longestCommonSubsequence1(num1, num2))
	println(longestCommonSubsequence2(num1, num2))
	println(longestCommonSubstring1("ABCD", "BCDA"))
}
