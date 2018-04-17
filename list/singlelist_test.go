package list

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	list := NewSingleLinkedList()
	fmt.Println("length:", list.size)
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	fmt.Println(list.GetAll())
	fmt.Println("length:", list.size)
	fmt.Println("GetByIndex 2nd:", list.GetByIndex(2).data)
	fmt.Println("insert(2,x)", list.Insert(2, "x"))
	fmt.Println("length:", list.size)
	fmt.Println(list.GetAll())
	fmt.Println("GetByIndex 2nd:", list.GetByIndex(2).data)
	fmt.Println("addFirst:", list.AddFirst("first"))
	fmt.Println(list.GetAll())
	fmt.Println("addLast:", list.AddLast("last"))
	fmt.Println(list.GetAll())
	fmt.Println("RemoveByIndex(3)", list.RemoveByIndex(3))
	fmt.Println(list.GetAll())
	fmt.Println("RemoveFirst:", list.RemoveFirst())
	fmt.Println("RemoveLast:", list.RemoveLast())
	fmt.Println(list.GetAll())
	fmt.Println("RemoveData(b):", list.RemoveData("b"))
	fmt.Println(list.GetAll())
	fmt.Println("GetFirst():", list.GetFirst().data)
	fmt.Println("GetLast():", list.GetLast().data)
}
