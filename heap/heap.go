package myheap

type MaxHeap struct {
	data     []int // 数组存储堆
	count    int   // 当前堆容量
	capacity int   //初始化容量
}

// 构造函数, 构造一个空堆, 可容纳capacity个元素
func NewMaxHeap(cap int) *MaxHeap {
	heap := new(MaxHeap)
	heap.data = make([]int, cap+1) //跳过0从1开始
	heap.count = 0
	heap.capacity = cap
	return heap
}

// Heapify：给定一个数组排列成 堆 的形状的过程
// 数组按照 二叉树 排列，每个叶子节点本身就是 最大堆
// 完全二叉树 第一个非叶子节点=最后一个索引/2，从后向前依次考察每个不是叶子节点的节点，然后shiftDown，继续向上
// O(n)
func NewMaxHeapHeapify(arr []int) *MaxHeap {
	n := len(arr)
	heap := new(MaxHeap)
	heap.data = make([]int, n+1) //跳过0从1开始
	heap.capacity = n
	for i := 0; i < n; i++ {
		heap.data[i+1] = arr[i]
	}
	heap.count = n
	for i := heap.count / 2; i >= 1; i-- { // 从第一个非叶子节点开始，叶子节点都是最大堆
		heap.shiftDown(i)
	}
	return heap
}

// 返回堆中的元素个数
func (h MaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值, 表示堆中是否为空
func (h MaxHeap) IsEmpty() bool {
	return h.count == 0
}

func (h *MaxHeap) Insert(item int) {
	if h.count+1 <= h.capacity {
		h.data[h.count+1] = item
		h.count++
		h.shiftUp(h.count)
	}
}

// 从最大堆中取出堆顶元素, 即堆中所存储的最大数据
func (h *MaxHeap) ExtractMax() int {
	var ret int
	if h.count > 0 {
		ret = h.data[1]
		h.data[1], h.data[h.count] = h.data[h.count], h.data[1]
		h.count--
		h.shiftDown(1)
	}
	return ret
}

// 新插入数据与父节点比较
func (h *MaxHeap) shiftUp(k int) {
	for k > 1 && h.data[k/2] < h.data[k] {
		h.data[k], h.data[k/2] = h.data[k/2], h.data[k]
		k /= 2
	}
}

// 将最后一个元素放置顶端，然后向下排序(谁大跟谁换)
func (h *MaxHeap) shiftDown(k int) {
	for 2*k <= h.count { // k存在左子节点
		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= h.count && h.data[j+1] > h.data[j] {
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if h.data[k] > h.data[j] {
			break
		}
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}
