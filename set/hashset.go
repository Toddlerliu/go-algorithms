package set

// 空对象 struct{}
// 可以和普通结构一样操作
// 不占用内存 unsafe.Sizeof(struct{})=0
type HashSet struct {
	items map[interface{}]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{
		items: make(map[interface{}]struct{}),
	}
}

func (set *HashSet) IsEmpty() bool {
	return len(set.items) == 0
}

func (set *HashSet) Size() int {
	return len(set.items)
}

func (set *HashSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}
	return true
}

func (set *HashSet) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

func (set *HashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

func (set *HashSet) Clear() {
	set.items = make(map[interface{}]struct{})
}
