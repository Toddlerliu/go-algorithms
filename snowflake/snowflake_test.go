package main

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	var i int64
	for i = 1; i <= 10; i++ {
		go func() {
			node, _ := NewIdWorker(i, i)
			fmt.Println(node.GenerateID())
		}()
	}
}
