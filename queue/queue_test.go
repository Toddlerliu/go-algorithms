package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Offer("a")
	queue.Offer("b")
	queue.Offer("c")
	fmt.Println("size:", queue.Size())
	fmt.Println("peek:", queue.Peek())
	fmt.Println("all:", queue.GetAll())
	v := queue.Poll()
	fmt.Println("poll: value", v)
	fmt.Println("all:", queue.GetAll())
	fmt.Println("size:", queue.Size())
}
