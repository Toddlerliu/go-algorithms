package list

type DNode struct {
	data interface{} // 存储的数据
	prev *DNode
	next *DNode
}

type DoubleList struct {
	size int
	head *DNode
	tail *DNode
}

func NewDoubleList(cap int)  {

}