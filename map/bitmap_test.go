package mmap

import (
	"testing"
	"fmt"
)

func TestBitMap(t *testing.T) {
	m := NewBitMap(10)
	m.SetBit(1)
	m.SetBit(3)
	m.SetBit(5)
	m.SetBit(7)
	fmt.Println("是否包含5？", m.Contain(5))
	fmt.Println(m.PrintBit())
	fmt.Println(m.PrintNum())
	m.Clear(5)
	fmt.Println("是否包含5？", m.Contain(5))
}
