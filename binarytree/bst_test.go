package mybinarysearch

import (
	"fmt"
	"testing"
)

func TestBSTInset(t *testing.T) {
	bst := NewBinarySearchTree()
	fmt.Println("空？", bst.IsEmpty())
	fmt.Println("inserting...")
	bst.Insert("a", "a")
	bst.Insert("b", "b")
	bst.Insert("c", "c")
	bst.Insert("d", "d")
	bst.Insert("e", "e")
	bst.Insert("f", "f")
	bst.Insert("g", "g")
	fmt.Println("空？", bst.IsEmpty())
	fmt.Println("e的value为：", bst.Get("e"))
	fmt.Println("max的key为：", bst.MaxKey(), "max的value为：", bst.MaxKeyValue())
	fmt.Println("min的key为：", bst.MinKey(), "min的value为：", bst.MinKeyValue())
}
