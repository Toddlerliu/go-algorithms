package queue

import (
	"testing"
	"fmt"
)

func TestCQueue(t *testing.T) {
	q, err := NewCQueue(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("queue current size is %d", q.Size())
	fmt.Println()
	fmt.Println("is empty :", q.IsEmpty())
	fmt.Println("is full :", q.IsFull())

	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	b, err := q.Offer(4)
	b, err = q.Offer(5)
	fmt.Println("入库情况：", b, err)
	fmt.Printf("queue current size is %d", q.Size())
	fmt.Println()
	fmt.Println("is empty :", q.IsEmpty())
	fmt.Println("is full :", q.IsFull())
	fmt.Println("head:", q.head, ":tail", q.tail)
	fmt.Println("queue:", q.GetAll())

	ele := q.Poll()
	fmt.Println("出队元素：", ele)
	ele = q.Poll()
	fmt.Println("出队元素：", ele)
	fmt.Println("queue:", q.GetAll())

	fmt.Println("插入5和6")
	q.Offer(5)
	q.Offer(6)
	fmt.Println("queue:", q.GetAll())

	fmt.Println("clear queue")
	q.Clear()
	fmt.Println("is empty :", q.IsEmpty())
	fmt.Println("is full :", q.IsFull())
	fmt.Println("queue:", q.GetAll())
	fmt.Println("head is :", q.head, " ;tail is :", q.tail)
}
