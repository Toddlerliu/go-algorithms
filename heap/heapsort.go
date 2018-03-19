package myheap

// 将所有的元素依次添加到堆中, 在将所有元素从堆中依次取出来, 即完成了排序
// 无论是创建堆的过程, 还是从堆中依次取出元素的过程, 时间复杂度均为O(nlogn)
// 整个堆排序的整体时间复杂度为O(nlogn)
func HeapSort1(arr []int) []int {
	n := len(arr)
	maxHeap := NewMaxHeap(n)
	// 将n个元素逐个插入到一个空堆中，O(nlogn)
	for  i:= 0; i< n; i++ {
		maxHeap.Insert(arr[i])
	}
	for i := n-1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
	return arr	
}

// heapify 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 实践复杂度为O(nlogn)
// 堆排序的总体时间复杂度依然是O(nlogn), 但是比HeapSort1性能更优, 因为创建堆的性能更优
func HeapSort2(arr []int) []int {
	maxHeap := NewMaxHeapHeapify(arr)
	for i := len(arr)-1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
	return arr
}

// 原地堆排序：不使用一个额外的最大堆, 直接在原数组上进行原地的堆排序
// 			  时间复杂度O(nlogn)；空间复杂度O(1)，无需额外数组空间
// 索引从0开始：parent(i) = (i-1)/2
//			   left child(i) = 2*i+1
// 			   right child(i) = 2*i+2 
// heapify shiftDown
func HeapSort(arr []int) []int {
	n := len(arr)
	
	// Heapify 堆从0开始索引，从(最后一个元素的索引-1)/2开始（第一个非叶子节点，叶子节点都是MaxHeap），最后一个元素的索引 = n-1
	for i := n-1; i >= 0; i-- {
		shiftDown3(arr,n,i) // 2
	}
	// 最大堆的最大元素放入 有序数组
	for i := n-1; i > 0; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		// 交换后不是MaxHeap，继续HeapDown
		shiftDown3(arr,i,0) // 2
	}
	return arr
}

// n：一共多少个元素；k：索引k下移
func shiftDown2(arr []int, n,k int)  {
	for 2*k+1 < n {
		j := 2*k+1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if arr[k] >= arr[j] {
			break			
		}

		arr[k],arr[j] = arr[j],arr[k]
		k = j
	}
}

// 优化:使用赋值的方式取代不断的交换
func shiftDown3(arr []int, n,k int)  {
	e := arr[k]
	for 2*k+1 < n {
		j := 2*k+1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if e >= arr[j] {
			break
		}

		arr[k] = arr[j]
		k = j
	}
	arr[k] = e
}