package list

import (
	"testing"
	"fmt"
)

func TestSingleList(t *testing.T) {
	list := NewSingleList()
	//fmt.Println(list.PrintSingleList())
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	//fmt.Println(list.PrintSingleList())
	fmt.Println("2st:",list.GetByIndex(2))
}
