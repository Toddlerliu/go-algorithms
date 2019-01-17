package mysort

import (
	"math/rand"
	"time"
)

// O(n^2)

// 选择排序：找到待排序列 最小值 放到最前位置（交换）,从下一位置继续查找最小值
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// 寻找[i,len(arr)) 区间里最小的值
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 冒泡排序:对相邻的元素进行两两比较，每一趟会将最大（小）的元素“浮”到最右端
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 鸡尾酒排序：冒泡排序完再返回到头
// O(n^2)
func CocktailSort(arr []int) []int {
	swap := true // 是否发生数据交换
	for swap {
		swap = false
		for i := 0; i < len(arr)/2; i++ {
			// 将最大值排到队尾
			for j := i; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swap = true
				}
			}
			// 将最小值排到队头
			for j := len(arr) - 1 - i - 1; j > i; j-- {
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
					swap = true
				}
			}
		}
	}
	return arr
}

// 奇偶排序：1、奇偶；2、偶奇...成对出现
// O(n^2)    多核：O(n^2/(m/2)) (m：待排个数)
func OddEvenSort(arr []int) []int {
	swap, oddOrEven := true, 0 // swap:是否发生交换；oddOrEven：0偶交换，1奇数交换
	for swap == true || oddOrEven == 1 {
		swap = false
		for i := oddOrEven; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
		if oddOrEven == 0 {
			oddOrEven = 1
		} else {
			oddOrEven = 0
		}
	}
	return arr
}

//// 并行奇偶排序
//func OddEvenSort2(arr []int) []int {
//	mu := sync.Mutex{}
//	swap, oddOrEven := true, 0
//	setSwap := func(b bool) {
//		mu.Lock()
//		swap = b
//		mu.Unlock()
//	}
//	getSwap := func() bool {
//		mu.Lock()
//		defer mu.Unlock()
//		return swap
//	}
//	for getSwap() == true || oddOrEven == 1 {
//		setSwap(false)
//		wg := &sync.WaitGroup{}
//		count := 0 // array.length/2 -(array.length%2 == 0?start:0)
//		if len(arr)%2 == 0 {
//			count = len(arr)/2 - oddOrEven
//		} else {
//			count = len(arr)/2 - 0
//		}
//		wg.Add(count)
//		for i := oddOrEven; i < len(arr)-1; i += 2 {
//			go func(i int, wg sync.WaitGroup) {
//				defer wg.Done()
//				if arr[i] > arr[i+1] {
//					arr[i], arr[i+1] = arr[i+1], arr[i]
//					setSwap(true)
//				}
//			}(i, *wg)
//		}
//		wg.Wait()
//		if oddOrEven == 0 {
//			oddOrEven = 1
//		} else {
//			oddOrEven = 0
//		}
//	}
//	return arr
//}

// 插入排序：依次选择待排序列，放入到 有序（已排序）列的合适位置(依次和前一个比较，交换)
// 待排序列有序效率极高 O(n)
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ { // 第0个元素有序
		e := arr[i] // copy
		var j int   // j保存元素e应该插入的位置
		// 元素arr[i]插入 有序（前面） 序列合适位置
		// for j = i; j > 0 && arr[j] < arr[j-1]; j-- { // 相比选择排序可以提前结束
		// 	arr[j],arr[j-1] = arr[j-1],arr[j] // 交换、赋值耗时（两两比较、交换）
		// }
		for j = i; j > 0 && e < arr[j-1]; j-- {
			arr[j] = arr[j-1] // 前面的往后挪一位
		}
		arr[j] = e
	}
	return arr
}

// 对arr[l...r]范围的数组进行插入排序
func InsertionSortPart(arr []int, l, r int) []int {
	for i := l + 1; i <= r; i++ { // 第0个元素有序
		e := arr[i] // copy
		var j int
		for j = i; j > l && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
	return arr
}

// 希尔排序(缩小增量法) ：加大插入排序中元素之间的间隔，并在这些有间隔的元素中进行插入排序
func ShellSort(arr []int) []int {
	n := len(arr)
	// 计算 increment sequence: 1, 4, 13, 40, 121, 364, 1093...
	// Knuth 增量序列:递推公式 h1=1, h(i) = 3 ∗ h(i−1) + 1
	h := 1 // 增量
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		// h-sort the array
		for i := h; i < n; i++ {

			e := arr[i] // copy
			var j = i
			// 对 arr[i], arr[i-h], arr[i-2*h], arr[i-3*h]... 使用插入排序
			for j = i; j >= h && e < arr[j-h]; j -= h {
				// arr[j],arr[j-h] = arr[j-h],arr[j]
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		// 计算下一个增量h
		h /= 3
	}
	return arr
}

// O(n*log n)

// 归并排序，自顶向下，逐步递归
func MergeSort(arr []int) []int {
	_mergeSort(arr, 0, len(arr)-1)
	return arr
}

// 递归使用归并排序，对arr[l...r]的范围进行排序,lr位置
func _mergeSort(arr []int, l, r int) {
	// if l >= r {
	// 	return
	// }

	// 优化：递归量很小的时候，数组近乎有序的概率较大，使用插入排序提高性能
	if r-l <= 15 { //16元素
		InsertionSortPart(arr, l, r)
		return
	}

	mid := (l + r) / 2 // 当l和r很大时，会发生溢出错误
	_mergeSort(arr, l, mid)
	_mergeSort(arr, mid+1, r)
	// 优化：有序则不需要merge
	if arr[mid] > arr[mid+1] {
		_merge(arr, l, mid, r)
	}
}

// 将arr[l...mid]和[mid+1...r]两部分进行归并
func _merge(arr []int, l, mid, r int) {
	tmp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		tmp[i-l] = arr[i] //tmp从零开始，arr从l开始，之间差l个偏移量
	}
	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i, j := l, mid+1
	for k := l; k <= r; k++ { //k最终排序的位置
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = tmp[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = tmp[i-l]
			i++
		} else if tmp[i-l] < tmp[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = tmp[i-l]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = tmp[j-l]
			j++
		}
	}
}

// 归并排序，自底向上，迭代，非递归
func MergeSortBottomUp(arr []int) []int {
	for size := 1; size <= len(arr); size += size { // 分组 1,2,4,8,18...
		for i := 0; i+size < len(arr); i += size + size { // 两个size区域merge；越界问题：i + size
			// 对arr[i...i+size-1]和arr[i+size...i+2*size-1]进行归并
			// 越界问题：数组末尾不足size个元素，i+size+size-1越界，可用n-1
			var minR int
			if i+size+size-1 < len(arr)-1 {
				minR = i + size + size - 1
			} else {
				minR = len(arr) - 1
			}
			_merge(arr, i, i+size-1, minR)
		}
	}
	return arr

	// // 优化
	// // 对于小数组, 使用插入排序优化
	// for i := 0 ; i < len(arr) ; i += 16 {
	// 	InsertionSort.sort(arr, i, Math.min(i+15, len(arr)-1))
	// }

	// for sz := 16; sz < n ; sz += sz {
	// 	for i := 0 ; i < n - sz ; i += sz+sz {
	// 		// 对于arr[mid] <= arr[mid+1]的情况,不进行merge
	//         if arr[i+sz-1] > arr[i+sz] {
	// 			_merge(arr, i, i+sz-1, Math.min(i+sz+sz-1,n-1) );
	// 		}
	// 	}
	// }
}

// 快速排序:
func QuickSort(arr []int) []int {
	_quickSort(arr, 0, len(arr)-1)
	return arr
}

// 对arr[l...r]部分快排
func _quickSort(arr []int, l, r int) {
	// if l >= r {
	// 	return
	// }
	// 优化：递归量很小的时候，数组近乎有序的概率较大，使用插入排序提高性能
	if r-l <= 15 { //16元素
		InsertionSortPart(arr, l, r)
		return
	}
	p := _partition(arr, l, r) // 分界处索引
	_quickSort(arr, l, p-1)
	_quickSort(arr, p+1, r)
}

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func _partition(arr []int, l, r int) int {
	// 优化：待排数组有序，快排效率最差O(n^2),初始值选择改进,随机化，随机算法，退化为O(n^2)的概率极低
	arr[l], arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l] = arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l], arr[l]

	v := arr[l] //基准
	// arr[l+1...j] < v ; arr[j+1...i) > v
	var j = l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1] // ++j
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 双路快速排序：
func QuickSort2Ways(arr []int) []int {
	_quickSort2(arr, 0, len(arr)-1)
	return arr
}

// 对arr[l...r]部分快排
func _quickSort2(arr []int, l, r int) {
	// if l >= r {
	// 	return
	// }
	// 优化：递归量很小的时候，数组近乎有序的概率较大，使用插入排序提高性能
	if r-l <= 15 { //16元素
		InsertionSortPart(arr, l, r)
		return
	}
	p := _partition2(arr, l, r) // 分界处索引
	_quickSort2(arr, l, p-1)
	_quickSort2(arr, p+1, r)
}

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func _partition2(arr []int, l, r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	// 优化：待排数组有序，快排效率最差O(n^2),初始值选择改进,随机化，随机算法，退化为O(n^2)的概率极低
	arr[l], arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l] = arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l], arr[l]

	v := arr[l] //基准

	// arr[l+1...i) <= v; arr(j...r] >= v
	i, j := l+1, r //临界点
	for {
		// 边界: 不能用<=,造成两棵子树不平衡
		// a. 对于arr[i]<v和arr[j]>v的方式，第一次partition得到的分点是数组中间；
		// b. 对于arr[i]<=v和arr[j]>=v的方式，第一次partition得到的分点是数组的倒数第二个。
		// 这是因为对于连续出现相等的情况，a方式会交换i和j的值；而b方式则会将连续出现的这些值归为其中一方，使得两棵子树不平衡
		for arr[i] < v && i <= r {
			i++
		}
		for arr[j] > v && j >= l+1 {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	// 循环结束后，i从前向后第一个>=v的位置；j从后向前第一个（整个数组最后一个）<=v的位置
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 三路快速排序O(nlogn)：重复键值效率极高
// 可以在1秒之内轻松处理100万数量级的数据
func QuickSort3Ways(arr []int) []int {
	_quickSort3(arr, 0, len(arr)-1)
	return arr
}

// 对arr[l...r]部分快排,分为<v, ==v, >v 三部分，递归对<v和>v两部分继续快排
func _quickSort3(arr []int, l, r int) {
	// if l >= r {
	// 	return
	// }
	// 优化：递归量很小的时候，数组近乎有序的概率较大，使用插入排序提高性能
	if r-l <= 15 { //16元素
		InsertionSortPart(arr, l, r)
		return
	}
	// partition
	arr[l], arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l] = arr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(r-l+1)+l], arr[l]
	v := arr[l]

	lt := l     // arr[l+1...lt] < v, 初始为空
	gt := r + 1 // arr[gt...r] > v, 初始为空
	i := l + 1  // arr[lt+1...i) == v, 初始为空
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l] // 这部操作未进行lt--
	// done：arr[l...lt-1]<v; arr[lt...gt-1]==v;arr[gt...r]>v
	_quickSort3(arr, l, lt-1)
	_quickSort3(arr, gt, r)
}

// 逆序对（归并排序求）
func InversePairs(arr []int) int64 {
	return inversePairs(arr, 0, len(arr)-1)
}

// 求arr[l..r]范围的逆序数对个数
func inversePairs(arr []int, l, r int) int64 {
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	// 求出 arr[l...mid] 范围的逆序数
	res1 := inversePairs(arr, l, mid)
	// 求出 arr[mid+1...r] 范围的逆序数
	res2 := inversePairs(arr, mid+1, r)
	return res1 + res2 + merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) int64 {
	tmp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		tmp[i-l] = arr[i] //tmp从零开始，arr从l开始，之间差l个偏移量
	}
	// 初始化逆序数对个数 res = 0
	var res int64
	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i, j := l, mid+1
	for k := l; k <= r; k++ { //k最终排序的位置
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = tmp[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = tmp[i-l]
			i++
		} else if tmp[i-l] < tmp[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = tmp[i-l]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = tmp[j-l]
			j++
			// 此时, 因为右半部分k所指的元素小
			// 这个元素和左半部分的所有未处理的元素都构成了逆序数对
			// 左半部分此时未处理的元素个数为 mid - j + 1
			res += (int64)(mid - i + 1)
		}
	}
	return res
}

// 求slice中第k大元素(快排)
// func SelectionMaxOfNth(arr []int, k int) int {
// 	if k<0 && k>len(arr) {
// 		fmt.Println("wrong params")
// 	}
// 	return selectionMaxOfNth()
// }
// func selectionMaxOfNth(arr []int, l,r,k int) int {

// }
