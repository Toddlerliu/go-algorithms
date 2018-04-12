package lru

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
	fmt.Println("front:", c.list.Front().Value.(*entry).key, c.list.Front().Value.(*entry).value)
	c.Get("a")
	fmt.Println("front:", c.list.Front().Value.(*entry).key, c.list.Front().Value.(*entry).value)
	c.Get(1)
	fmt.Println("front:", c.list.Front().Value.(*entry).key, c.list.Front().Value.(*entry).value)

	fmt.Println("last:", c.list.Back().Value.(*entry).key, c.list.Back().Value.(*entry).value)

	fmt.Println("remove oldest:")
	c.RemoveOldest()
	fmt.Println("last:", c.list.Back().Value.(*entry).key, c.list.Back().Value.(*entry).value)
	fmt.Println("remove :")
	c.Remove(3)
	fmt.Println("last:", c.list.Back().Value.(*entry).key, c.list.Back().Value.(*entry).value)
	fmt.Println("last:", c.cache[3])
}
