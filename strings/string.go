package strings

// 寻找最长不含重复字符的子串
func lengthOfNonRepeatingSubStr(str string) int {
	// 某字符最后出现的位置
	// lastOccured[x]不存在，或 <start ，无需操作
	// lastOccured[x]>=start ，更新start
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	// 更新lastOccured[x]和maxLength
	for i, ch := range []rune(str) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
