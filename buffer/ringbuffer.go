package buffer

import (
	"errors"
)

// Disruptor
type RingBuffer struct {
	data  []byte
	size  int64 // 总容量
	head  int64 // 头指针, 与环形队列不同的是，没有尾指针，不删除数据，直到新的数据覆盖他们
	count int64 // 写入的容量，标识是否跨环
}

func NewRingBuffer(size int64) (*RingBuffer, error) {
	if size <= 0 {
		return nil, errors.New("error size")
	}
	buffer := &RingBuffer{
		size: size,
		data: make([]byte, size),
	}
	return buffer, nil
}

func (b RingBuffer) Size() int64 {
	return b.count
}

func (b *RingBuffer) Write(ele []byte) {

	n := len(ele)
	b.count += int64(n)

	// 数量大于容量，截取最新
	if int64(n) > b.size {
		ele = ele[int64(n)-b.size:]
	}
	copy(b.data[b.head:], ele)

	remain := b.size - b.head
	if int64(len(ele)) > remain {
		// 跨环
		copy(b.data, ele[remain:])
	}

	b.head = (b.head + int64(len(ele))) % b.size
}

func (b RingBuffer) Read(index int64) byte {
	index = index % b.size
	return b.data[index]
}

func (b RingBuffer) ReadAll() []byte {
	switch {
	case b.count >= b.size && b.head == 0: // 整个环
		return b.data
	case b.count > b.size: // 跨环
		ret := make([]byte, b.size)
		copy(ret, b.data[b.head:])
		copy(ret[b.size-b.head:], b.data[:b.head])
		return ret
	default: // 一个环内未满
		return b.data[:b.head]
	}
}
