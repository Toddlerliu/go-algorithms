package list

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	list := NewSkipList()
	for i := 0; i < 50; i++ {
		list.Insert(i, nil)
	}
	fmt.Println("是否包含5？：", list.Contains(5))
	fmt.Println("是否包含11？：", list.Contains(11))
	list.levelPrint()
	list.delete(11)
	list.levelPrint()
}
