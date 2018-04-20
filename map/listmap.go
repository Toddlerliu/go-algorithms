package mmap

type Pair struct {
	Key   interface{}
	Value interface{}
}

// 有序map(插入顺序)
type ListMap struct {
	keys  []interface{}
	pairs map[interface{}]*Pair
}

func NewListMap() *ListMap {
	return &ListMap{
		keys:  make([]interface{}, 0),
		pairs: make(map[interface{}]*Pair),
	}
}

func (m *ListMap) Size() int {
	return len(m.keys)
}

func (m *ListMap) Contain(k interface{}) bool {
	if len(m.pairs) == 0 {
		return false
	}
	_, ok := m.pairs[k]
	return ok
}

// 如果key存在，更新；不存在，新增；插入顺序
func (m *ListMap) Add(k, v interface{}) {
	if value, ok := m.pairs[k]; ok {
		value.Value = v
	} else {
		m.keys = append(m.keys, k)
		m.pairs[k] = &Pair{k, v}
	}
}

func (m *ListMap) Get(k interface{}) (interface{}, bool) {
	if m.Contain(k) {
		return m.pairs[k].Value, true
	}
	return nil, false
}

func (m *ListMap) GetTop() (*Pair, bool) {
	if len(m.keys) > 0 {
		v, ok := m.pairs[m.keys[0]]
		return v, ok
	}
	return nil, false
}

func (m *ListMap) Keys() []interface{} {
	return m.keys
}

func (m *ListMap) Values() []interface{} {
	res := make([]interface{}, 0)
	for _, v := range m.keys {
		if value, ok := m.pairs[v]; ok {
			res = append(res, value.Value)
		}
	}
	return res
}

func (m *ListMap) Pairs() map[interface{}]*Pair {
	return m.pairs
}

func (m *ListMap) Remove(k interface{}) bool {
	if m.Contain(k) {
		delete(m.pairs, k)
		for i, v := range m.keys {
			if v == k {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
	return false
}

func (m *ListMap) RemoveTop() bool {
	if len(m.keys) > 0 {
		return m.Remove(m.keys[0])
	}
	return false
}

func (m *ListMap) Iteraror() <-chan *Pair {
	ch := make(chan *Pair, len(m.keys))
	go func() {
		defer close(ch)
		for _, v := range m.keys {
			if value, ok := m.pairs[v]; ok {
				ch <- value
			}
		}
	}()
	return ch
}
