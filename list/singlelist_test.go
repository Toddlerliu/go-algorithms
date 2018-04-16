package list

import (
	"testing"
	"fmt"
)

func TestSingleList(t *testing.T) {
	list := NewSingleList()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add(1)
	fmt.Println(list.GetAll())
	fmt.Println("length:",list.size)
	fmt.Println("2nd:",list.GetByIndex(2))
}
