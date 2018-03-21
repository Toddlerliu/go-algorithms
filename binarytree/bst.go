package mybinarysearch

type Node struct {
	Key string //k-v都是string
	Value string
	Left *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node	
	n int
}

func (tree BinarySearchTree) Insert (value string) {

}