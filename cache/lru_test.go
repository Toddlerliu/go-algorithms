package cache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(5)
	fmt.Println("len:", c.Size())
	c.Set(1, 1)
	c.Set(2, 2)
	c.Set(3, 3)
	c.Set("a", "a")
	c.Set("b", "b")
	fmt.Println("len:", c.Size())
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("get a")
	c.Get("a")
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("get 1")
	c.Get(1)
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("remove oldest:")
	c.RemoveOldest()
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("remove :")
	c.Remove(3)
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
