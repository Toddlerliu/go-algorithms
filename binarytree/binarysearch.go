package mybinarysearch

// 二分查找：有序 数组arr中,查找target
// 			找到返回true和相应索引index
//  		未找到返回false和-1
// O(logn)
func BinarySearch(arr []string, target string) (bool, int) {
	// arr[l...r]中查找target
	l, r := 0, len(arr)-1
	for l <= r {
		//mid := (l + r)/2 // 两个int相加若超出int最大值溢出
		mid := l + (r-l)/2

		if arr[mid] == target {
			return true, mid
		}
		if arr[mid] > target {
			// arr[l...mid-1]
			r = mid - 1
		} else {
			// arr[mid+1...r]
			l = mid + 1
		}
	}
	return false, -1
}

// 递归二分查找
func BinarySearchRecursion(arr []string, target string) (bool, int) {
	return binarySearchRecursion(arr, 0, len(arr)-1, target)
}

func binarySearchRecursion(arr []string, l, r int, target string) (bool, int) {
	if l > r {
		return false, -1
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return true, mid
	} else if arr[mid] > target {
		return binarySearchRecursion(arr, 0, mid-1, target)
	} else {
		return binarySearchRecursion(arr, mid+1, r, target)
	}
}

// 伪二分查找
func FBinarySearch(arr []string, target string) (bool, int) {
	l, r := 0, len(arr)-1
	if r < 0 {
		return false, -1
	}
	for l <= r {
		if arr[l] == target {
			return true, l
		}
		if arr[r] == target {
			return true, r
		}
		l++
		r--
	}
	return false, -1
}

// 二分查找变种：floor、ceil
// 二分查找存在问题：重复值下无法确定返回的唯一索引值
// floor：是查找元素在数组中第一个索引位置；若数组中无此元素，则是target的前一个索引位置。
// ceil：是查找元素在数组中最后一个索引位置；若数组中无此元素，则是target的后一个索引位置。

// 如果找到target, 返回第一个target的索引index
// 如果没有找到target, 返回比target小的最大值相应的索引, 如果这个最大值有多个, 返回最大索引
// 如果这个target比整个数组的最小元素值还要小, 则不存在这个target的floor值, 返回-1
func Floor(arr []string, target string) (bool, int) {
	if target < arr[0] {
		return false, -1
	}
	l, r := -1, len(arr)-1
	for l < r {
		// 寻找比target小的最大索引
		mid := l + (r-l+1)/2
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if l == r {
		// 如果该索引+1就是target本身, 该索引+1即为返回值
		if l+1 < len(arr) && arr[l+1] == target {
			return true, l + 1
		}
		return false, l
	}
	return false, -999
}

// 返回重复target最后一个索引，如无，返回target前元素的后一个索引
func Ceil(arr []string, target string) (bool, int) {
	if target > arr[len(arr)-1] { //比最大还要大
		return false, -1
	}
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == r {
		// 如果该索引-1就是target本身, 该索引+1即为返回值
		if r-1 >= 0 && arr[r-1] == target {
			return true, r - 1
		}
		return false, r
	}
	return false, -999
}

// 			查找元素	插入元素	删除元素
// 普通数组	  O(n)		 O(n)		 O(n)
// 顺序数组  O(logn)	 O(n)		 O(n)
// 二分搜索树O(logn)    O(logn)     O(logn)
