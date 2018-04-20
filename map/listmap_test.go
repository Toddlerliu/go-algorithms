package mmap

import (
	"fmt"
	"testing"
)

func TestListMap(t *testing.T) {
	lm := NewListMap()
	fmt.Println("Size:", lm.Size())
	fmt.Println("Inserting...")
	lm.Add(1, 1)
	lm.Add(2, 2)
	lm.Add("a", "y")
	lm.Add("b", "b")
	lm.Add("b", "x")
	fmt.Println("Size:", lm.Size())
	fmt.Println("Contains(1)", lm.Contain(1))
	fmt.Println("Contains(c)", lm.Contain("c"))
	v, ok := lm.Get("b")
	fmt.Println("Get(b)", v, ok)
	v1, ok1 := lm.GetTop()
	fmt.Println("GetTop:", v1, ok1)
	fmt.Println("Keys:", lm.Keys())
	fmt.Println("Values:", lm.Values())
	fmt.Println("Pairs:", lm.Pairs())
	lm.Remove("b")
	lm.RemoveTop()
	fmt.Println("Keys:", lm.Keys())
	fmt.Println("Values:", lm.Values())
	vv := lm.Iteraror()
	fmt.Println("Pairs:", (<-vv).Key, (<-vv).Value)
}
