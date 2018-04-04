package main

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	node,_:=NewWorkerNode(10)
	fmt.Println(node.GenerateID())
}
