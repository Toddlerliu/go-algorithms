package myheap

import (
	"time"
	"math/rand"
	"log"
)

// 生成n个元素的随机数组，范围[l,r]
func GenerateRandomArray(n,l,r int) []int {
	if l >= r {
		log.Fatal("wrong number")
	}
	arr := make([]int, n) // var arr []int
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<n;i++{
		arr[i] = rand.Intn(r - l + 1) + l
	}
	return arr
}

// 检验是否排序
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return false
		}	
	}
	return true
}

// 性能测试
func TestSort() {
	
}