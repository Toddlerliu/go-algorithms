package mysort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", SelectionSort(arr))
}

func TestInsertionSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", InsertionSort(arr))
}

func TestBubbleSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", BubbleSort(arr))
}

func TestShellSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", ShellSort(arr))
}

func TestMergeSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", MergeSort(arr))
}

func TestMergeSortBottomUp(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10)
	fmt.Println("ori:", arr)
	fmt.Println("new:", MergeSortBottomUp(arr))
}

func TestQuickSort(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10) //大量相同
	// fmt.Println("ori:", arr)
	fmt.Println("快速排序:", QuickSort(arr))
	fmt.Println("双路快排:", QuickSort2Ways(arr))
	fmt.Println("三路快排:", QuickSort3Ways(arr))
}

func TestInversePairs(t *testing.T) {
	arr := GenerateRandomArray(10, 1, 10) //大量相同
	// fmt.Println("ori:", arr)
	fmt.Println("逆序对:", InversePairs(arr))
	InsertionSort(arr)
	fmt.Println("排序后逆序对:", InversePairs(arr))
}
