package set

import "container/list"

type LinkedHashSet struct {
	items map[interface{}]struct{}
	list  *list.List
}

func NewLinkedHashSet() *LinkedHashSet {
	return &LinkedHashSet{
		items: make(map[interface{}]struct{}),
		list:  list.New(),
	}
}

func (set *LinkedHashSet) IsEmpty() bool {
	return set.list.Len() == 0
}

func (set *LinkedHashSet) Size() int {
	return set.list.Len()
}

func (set *LinkedHashSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}
	return true
}

func (set *LinkedHashSet) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = struct{}{}
		set.list.PushBack(item)
	}
}

func (set *LinkedHashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
		for e := set.list.Front(); e != nil; e = e.Next() {
			if e.Value == item {
				set.list.Remove(e)
			}
		}
	}
}

func (set *LinkedHashSet) Clear() {
	set.items = make(map[interface{}]struct{})
	set.list.Init()
}
