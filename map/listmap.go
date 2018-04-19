package mmap

type Pair struct {
	Key   interface{}
	Value interface{}
}

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
