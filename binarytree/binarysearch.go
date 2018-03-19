
package mybinarysearch

// 二分查找：有序数组arr中,查找target
// 			找到返回true和相应索引index
//  		未找到返回false和-1
func BinarySearch(arr []Sortable, target Sortable) (bool, int) {
	// 在arr[l...r]之中查找target
	l,r := 0, arr.Len()-1
	for l <= r {
		//mid := (l + r)/2
		mid := l + (r-l)/2
		
		if arr[mid] == target {
			return true,mid
		}
		if arr[mid] > target {
			r = mid - 1
		}else {
			l = mid + 1
		}
	}
	return false,-1
}