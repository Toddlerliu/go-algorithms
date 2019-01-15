package mybinarysearch

import (
	"testing"
	"fmt"
)

func TestBinaryIndexedTree(t *testing.T) {
	tree := NewBinaryIndexedTree(9)
	for i := 1; i <= 8; i++ {
		tree.AddOrUpdate(i, i)
	}
	fmt.Println(tree.a)
	fmt.Printf("tree的前%d项和为：%d\n", 8, tree.Sum(8))
}
