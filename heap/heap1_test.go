package myheap

import (
	"testing"
	"container/heap"
	"fmt"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 3)
	heap.Push(h, 7)
	heap.Push(h, 6)
	fmt.Println(h)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
