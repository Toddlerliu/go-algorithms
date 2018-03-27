package mybinarysearch

import (
	"fmt"
	"testing"
)

func TestBSTInset(t *testing.T) {
	bst := NewBinarySearchTree()
	fmt.Println("空？", bst.IsEmpty())
	fmt.Println("inserting...")
	bst.Insert("28", "a")
	bst.Insert("16", "b")
	bst.Insert("30", "c")
	bst.Insert("13", "d")
	bst.Insert("22", "e")
	bst.Insert("29", "f")
	bst.Insert("42", "g")
	fmt.Println("空？", bst.IsEmpty())
	fmt.Println("max的key为：", bst.MaxKey(), "max的value为：", bst.MaxKeyValue())
	fmt.Println("min的key为：", bst.MinKey(), "min的value为：", bst.MinKeyValue())
	fmt.Println(bst.Search("100"))
	fmt.Println(bst.Search("30"))
	fmt.Println("前序排序：",bst.PreOrder())
	fmt.Println("中序排序：",bst.InOrder())
	fmt.Println("后序排序：",bst.PostOrder())
	fmt.Println("层序排序：",bst.LevelOrder())
	fmt.Println("删除最小：")
	bst.RemoveMin()
	fmt.Println("中序第一个为最小：",bst.InOrder())
	fmt.Println("删除最大：")
	bst.RemoveMax()
	fmt.Println("中序最后一个为最大：",bst.InOrder())
}
