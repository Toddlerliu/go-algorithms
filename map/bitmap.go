package mmap

import "fmt"

type BitMap struct {
	data []byte
	cap  uint64 // 容量
}

func NewBitMap(cap uint64) *BitMap {
	return &BitMap{
		data: make([]byte, (cap>>3)+1),
		cap:  cap,
	}
}

func (m *BitMap) SetBit(num uint64) {
	index := num >> 3 // num/8
	// %模运算： k % m = k & (m-1) (m=2的n次幂)
	position := num & 0x07            // num%8
	m.data[index] |= 0x01 << position // 置1
}

func (m *BitMap) Clear(num uint64) {
	index := num >> 3
	position := num & 0x07
	m.data[index] &^= 0x01 << position // &^ (and not): 与运算符左边数据相异的位保留，相同位清零
	// m.data[index] &= ^(0x01 << position)
}

func (m *BitMap) Contain(num uint64) bool {
	index := num >> 3
	position := num & 0x07
	return m.data[index]&(0x01<<position) != 0
}

func (m *BitMap) PrintBit() string {
	s := make([]uint64, m.cap)
	for i := uint64(0); i < m.cap; i++ {
		if m.Contain(i) {
			s[i] = 1
		}
	}
	return fmt.Sprintf("%v", s)
}

func (m *BitMap) PrintNum() string {
	s := make([]uint64, 0)
	for i := uint64(0); i < m.cap; i++ {
		if m.Contain(i) {
			s = append(s, i)
		}
	}
	return fmt.Sprintf("%v", s)
}
