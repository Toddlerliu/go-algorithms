package list

type SNode struct {
	data interface{} // 存储的数据
	next *SNode
}

type SingleList struct {
	size int
	head *SNode
	tail *SNode
}

func NewSingleList(cap int) *SingleList {
	return &SingleList{0,nil,nil}
}

func (l *SingleList) Add() {

}
