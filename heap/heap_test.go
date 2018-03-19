package myheap

import(
	"testing"
	"fmt"
)

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap(100)
	fmt.Println(heap.count)
	for i := 0; i < 15; i++ {
		heap.Insert(i)
	}
	fmt.Println(heap.data)
	for heap.count > 0 {
		fmt.Println(heap.ExtractMax())
	}
}

func TestHeapSort(t *testing.T)  {
	arr := GenerateRandomArray(10,1,10)
	fmt.Println("ori:", arr)
	fmt.Println("heapsort1:", HeapSort1(arr))
	fmt.Println("heapsort2:", HeapSort2(arr))
	fmt.Println("heapsort:", HeapSort(arr))
}

func TestNewIndexMaxHeap(t *testing.T) {
	fmt.Println("============Test IndexMaxHeap====")
	heap := NewIndexMaxHeap(6)
	heap.Insert(0,2)
	heap.Insert(1,10)
	heap.Insert(2,3)
	heap.Insert(3,4)
	heap.Update(3,11)
	fmt.Println("index:",heap.indexes)
	fmt.Println("data:",heap.data)
	fmt.Println("reverse:",heap.reverse)
	for i := 0; i < heap.capacity; i++ {
		fmt.Println(heap.ExtractMax())
		fmt.Println("index:",heap.indexes)
		fmt.Println("data:",heap.data)
		fmt.Println("reverse:",heap.reverse)
	}
}

func TestIndexMaxHeapHeapify(t *testing.T) {
	fmt.Println("============Test IndexMaxHeap=====array")
	arr := GenerateRandomArray(10,1,10)
	fmt.Println("ori:",arr)
	// index和reverse没值。×
	heap := NewIndexMaxHeapHeapify(arr)
	fmt.Println(heap.count)
	fmt.Println("index:",heap.indexes)
	fmt.Println("data:",heap.data)
	fmt.Println("reverse:",heap.reverse)
	
	for heap.count > 0 {
		fmt.Println(heap.ExtractMax())
	}
}