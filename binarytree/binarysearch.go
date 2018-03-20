
package mybinarysearch

// 二分查找：有序 数组arr中,查找target
// 			找到返回true和相应索引index
//  		未找到返回false和-1
// O(logn)
func BinarySearch(arr []string, target string) (bool, int) {
	// arr[l...r]中查找target
	l,r := 0, len(arr)-1
	for l <= r {
		//mid := (l + r)/2 // 两个int相加若超出int最大值溢出
		mid := l + (r-l)/2
		
		if arr[mid] == target {
			return true,mid
		}
		if arr[mid] > target { 
			// arr[l...mid-1]
			r = mid - 1
		} else { 
			// arr[mid+1...r]
			l = mid + 1
		}
	}
	return false,-1
}

// 递归二分查找
func BinarySearchRecursion(arr []string, target string) (bool,int) {
	return binarySearchRecursion(arr, 0, len(arr)-1, target)
}

func binarySearchRecursion(arr []string, l , r int ,target string) (bool,int) {
	if l>r {
		return false,-1
	}
	mid := (l + r)/2
	if arr[mid] == target {
		return true,mid
	} else if arr[mid] > target {
		return binarySearchRecursion(arr,0,mid-1,target)
	} else {
		return binarySearchRecursion(arr,mid+1,r,target)
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

// 			查找元素	插入元素	删除元素
// 普通数组	  O(n)		 O(n)		 O(n)
// 顺序数组  O(logn)	 O(n)		 O(n)
// 二分搜索树O(logn)    O(logn)     O(logn)

