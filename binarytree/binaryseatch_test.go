package mybinarysearch

import(
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println("======test BinarySearch")
	s := []string{"a","b","c","d","e","gg"}
	exist,index := BinarySearch(s,"d")
	if exist {
		fmt.Println("exist,index is:",index)
	}else {
		fmt.Println("not exist")
	}
}

func TestBinarySearchRecursion(t *testing.T) {
	fmt.Println("======test BinarySearchRecursion")
	s := []string{"a","b","c","d","e","gg"}
	exist,index := BinarySearchRecursion(s,"d")
	if exist {
		fmt.Println("exist,index is:",index)
	}else {
		fmt.Println("not exist")
	}
}

func TestFBinarySearch(t *testing.T) {
	fmt.Println("======test FBinarySearchs")
	s := []string{"a","b","c","d","e","gg"}
	exist,index := FBinarySearch(s,"z")
	if exist {
		fmt.Println("exist,index is:",index)
	}else {
		fmt.Println("not exist")
	}
}