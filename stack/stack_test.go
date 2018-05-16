package stack

import (
	"fmt"
	"testing"
)

func TestStacke(t *testing.T) {
	stack := NewStacke()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	fmt.Println("size:", stack.Size())
	fmt.Println("peek:", stack.Peek())
	fmt.Println("all:", stack.GetAll())
	v, err := stack.Pop()
	fmt.Println("pop: value", v, "err:", err)
	fmt.Println("all:", stack.GetAll())
	fmt.Println("size:", stack.Size())
}
