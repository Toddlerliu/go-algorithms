package mybinarysearch

import (
	"testing"
	"fmt"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	data := []string{"北京大学", "北京航空航天大学", "北京邮电大学", "北京师范大学", "北京科技大学", "北京理工大学", "北京交通大学", "北京大学国际法学院", "北京交通运输职业学院",
		"北京工业大学通州分校", "北京工业大学耿丹学院", "北京交通职业技术学院"}
	for _, value := range data {
		trie.Insert(value, value)
	}
	nodes := trie.PrefixSearch("北京交通", 100)
	for _, value := range nodes {
		fmt.Println(value.data.(string))
	}
}
