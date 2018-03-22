package myheap

import (
	"fmt"
)

type IndexMaxHeap struct {
	data     []int // 最大索引堆中的数据,类型可以变
	indexes  []int // 最大索引堆中的索引, indexes[x] = i 表示索引i在x的位置
	reverse  []int // 最大索引堆中的反向索引, reverse[i] = x 表示索引i在x的位置
	count    int   // 当前堆容量
	capacity int   //初始化容量
}

// 构造函数, 构造一个空堆, 可容纳capacity个元素
func NewIndexMaxHeap(cap int) *IndexMaxHeap {
	heap := new(IndexMaxHeap)
	heap.data = make([]int, cap+1) //跳过0从1开始
	heap.indexes = make([]int, cap+1)
	heap.reverse = make([]int, cap+1)
	for i := 0; i <= heap.capacity; i++ { //索引i在堆中的位置，初始化时没有，为0
		heap.reverse[i] = 0 // 从1开始，0表示跟本不存在
	}
	heap.count = 0
	heap.capacity = cap
	return heap
}

// Heapify：给定一个数组排列成 堆 的形状的过程
// 数组按照 二叉树 排列，每个叶子节点本身就是 最大堆
// 完全二叉树 第一个非叶子节点=最后一个索引/2，从后向前依次考察每个不是叶子节点的节点，然后shiftDown，继续向上
// O(n)
func NewIndexMaxHeapHeapify(arr []int) *IndexMaxHeap {
	n := len(arr)
	heap := new(IndexMaxHeap)
	heap.data = make([]int, n+1) //跳过0从1开始
	heap.indexes = make([]int, n+1)
	heap.reverse = make([]int, n+1)
	heap.capacity = n
	// for i := 0; i < n; i++ {
	// 	heap.data[i+1] = arr[i]
	// }
	copy(heap.data[1:], arr)
	heap.count = n
	for i := heap.count / 2; i >= 1; i-- { // 从第一个非叶子节点开始，叶子节点都是最大堆
		heap.shiftDown(i)
	}
	return heap
}

// 返回堆中的元素个数
func (h IndexMaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值, 表示堆中是否为空
func (h IndexMaxHeap) IsEmpty() bool {
	return h.count == 0
}

// 看索引i所在的位置是否存在元素
// 越界问题：i索引的元素真的存在在堆中，i在容量范围里不意味着一定在堆中
func (h IndexMaxHeap) Contain(i int) bool {
	if i+1 >= 1 && i+1 <= h.capacity {
		return h.reverse[i+1] != 0
	}
	return false
}

// 新元素的索引为i, 元素为item
// 传入的i对用户而言,是从0索引的,内部从1开始
func (h *IndexMaxHeap) Insert(i int, item int) {
	if h.count+1 <= h.capacity && i+1 >= 1 && i+1 <= h.capacity {
		// 再插入一个新元素前,还需要保证索引i所在的位置是没有元素的
		if !h.Contain(i) {
			i += 1
			h.data[i] = item
			h.indexes[h.count+1] = i
			h.reverse[i] = h.count + 1

			h.count++
			h.shiftUp(h.count)
		}
	}
}

// 从最大索引堆中取出堆顶元素, 即索引堆中所存储的最大数据
func (h *IndexMaxHeap) ExtractMax() int {
	var ret int
	if h.count > 0 {
		ret = h.data[h.indexes[1]]
		h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
		h.reverse[h.indexes[1]] = 1       // 第一个位置
		h.reverse[h.indexes[h.count]] = 0 //删除置0
		h.count--
		h.shiftDown(1)
	}
	return ret
}

// 从最大索引堆中取出堆顶元素的索引
// 传入的i对用户而言,是从0索引的
func (h *IndexMaxHeap) ExtractMaxIndex() int {
	var ret int
	if h.count > 0 {
		ret = h.indexes[1] - 1 // 1→0
		h.indexes[1], h.indexes[h.count] = h.indexes[h.count], h.indexes[1]
		h.reverse[h.indexes[1]] = 1       // 第一个位置
		h.reverse[h.indexes[h.count]] = 0 //删除置0
		h.count--
		h.shiftDown(1)
	}
	return ret
}

// 获取最大索引堆中索引为i的元素
func (h *IndexMaxHeap) GetItem(i int) int {
	if h.Contain(i) {
		return h.data[i+1]
	} else {
		fmt.Println("索引越界")
		return 0
	}
}

// 获取最大索引堆中的堆顶元素
func (h IndexMaxHeap) getMax() int {
	if h.count > 0 {
		return h.data[h.indexes[1]]
	}
	return 0
}

// 获取最大索引堆中的堆顶元素
func (h IndexMaxHeap) getMaxIndex() int {
	if h.count > 0 {
		return h.indexes[1] - 1
	}
	return 0
}

// 将最大索引堆中索引为i的元素修改为newItem
// O(n+nlogn) → O(n)
func (h *IndexMaxHeap) Update(i int, newItem int) {
	// 越界问题：i索引的元素真的存在在堆中，i在容量范围里不意味着一定在堆中
	if h.Contain(i) {
		i += 1
		h.data[i] = newItem

		// 找到indexes[j] = i, j表示data[i]在堆中的位置
		// 之后shiftUp(j), 再shiftDown(j)
		// for j := 1; j <= h.count; j ++ { // O(n)
		// 	if h.indexes[j] == i {
		// 		h.shiftUp(j) // O(nlogn)
		// 		h.shiftDown(j) // O(nlogn)
		// 		return
		// 	}
		// }

		// 有了 reverse 之后,
		// 我们可以非常简单的通过reverse直接定位索引i在indexes中的位置
		// 整体变为O(logn)
		j := h.reverse[i] // O(1)
		h.shiftUp(j)      // O(logn)
		h.shiftDown(j)    // O(logn)
	} else {
		fmt.Println("索引越界")
	}
}

// 新插入数据与父节点比较，k是数据的索引
func (h *IndexMaxHeap) shiftUp(k int) {
	for k > 1 && h.data[h.indexes[k/2]] < h.data[h.indexes[k]] {
		h.indexes[k], h.indexes[k/2] = h.indexes[k/2], h.indexes[k] // 交换索引，非数据
		h.reverse[h.indexes[k/2]] = k / 2
		h.reverse[h.indexes[k]] = k
		k /= 2
	}
}

// 将最后一个元素放置顶端，然后向下排序(谁大跟谁换)
func (h *IndexMaxHeap) shiftDown(k int) {
	for 2*k <= h.count { // k存在左子节点
		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= h.count && h.data[h.indexes[j+1]] > h.data[h.indexes[j]] {
			j += 1
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if h.data[h.indexes[k]] >= h.data[h.indexes[j]] {
			break
		}
		h.indexes[k], h.indexes[j] = h.indexes[j], h.indexes[k]
		h.reverse[h.indexes[k]] = k
		h.reverse[h.indexes[j]] = j
		k = j
	}
}
