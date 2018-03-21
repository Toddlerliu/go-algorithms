package mybinarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println("======test BinarySearch")
	s := []string{"a", "b", "c", "d", "e", "gg"}
	exist, index := BinarySearch(s, "d")
	if exist {
		fmt.Println("exist,index is:", index)
	} else {
		fmt.Println("not exist")
	}
}

func TestBinarySearchRecursion(t *testing.T) {
	fmt.Println("======test BinarySearchRecursion")
	s := []string{"a", "b", "c", "d", "e", "gg"}
	exist, index := BinarySearchRecursion(s, "d")
	if exist {
		fmt.Println("exist,index is:", index)
	} else {
		fmt.Println("not exist")
	}
}

func TestFBinarySearch(t *testing.T) {
	fmt.Println("======test FBinarySearchs")
	s := []string{"a", "b", "c", "d", "e", "gg"}
	exist, index := FBinarySearch(s, "z")
	if exist {
		fmt.Println("exist,index is:", index)
	} else {
		fmt.Println("not exist")
	}
}

func TestFloor(t *testing.T) {
	fmt.Println("======test Floor")
	s := []string{"b", "c", "gg","gg","gg","gg","x"}
	exist, index := Floor(s, "gg")
	if exist {
		fmt.Println("exist,index is:", exist,index)
	} else {
		fmt.Println("not exist",exist,index)
	}
}

func TestCeil(t *testing.T) {
	fmt.Println("======test Ceil")
	s := []string{"b", "c", "gg","gg","gg","gg","x"}
	exist, index := Ceil(s, "f")
	if exist {
		fmt.Println("exist,index is:", exist,index)
	} else {
		fmt.Println("not exist",exist,index)
	}
}
