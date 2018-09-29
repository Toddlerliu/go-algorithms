package buffer

import (
	"testing"
	"fmt"
)

func TestRingBuffer(t *testing.T) {
	buf, _ := NewRingBuffer(8)
	buf.Write([]byte("abcdefghi"))
	fmt.Println(string(buf.ReadAll()))
	fmt.Printf("index of 3 is %s", string(buf.Read(3)))
}
