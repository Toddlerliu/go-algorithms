// 泛型辅助
package mysort

// 参考sort源码实现go泛型
type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

// int型切片
type IntArr []int

func (arr IntArr) Len() int {
	return len(arr)
}

func (arr IntArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}

func (arr IntArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// folat类型切片
type floatArr []float32

func (arr floatArr) Len() int {
	return len(arr)
}

func (arr floatArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}

func (arr floatArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// string类型slice，按照字符串长度排序
type StrArr []string

func (arr StrArr) Len() int {
	return len(arr)
}

func (arr StrArr) Less(i int, j int) bool {
	return len(arr[i]) < len(arr[j])
}

func (arr StrArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
