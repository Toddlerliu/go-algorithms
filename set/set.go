package set

// 空对象 struct{}
// 可以和普通结构一样操作
// 不占用内存 unsafe.Sizeof(struct{})=0
type Set struct {
	items map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[interface{}]struct{}),
	}
}

func (set *Set) IsEmpty() bool {
	return len(set.items) == 0
}

func (set *Set) Size() int {
	return len(set.items)
}

func (set *Set) Contains(item interface{}) bool {
	if _, contains := set.items[item]; !contains {
		return false
	}
	return true
}

func (set *Set) Add(item interface{}) {
	set.items[item] = struct{}{}
}

func (set *Set) Remove(item interface{}) {
	delete(set.items, item)
}
