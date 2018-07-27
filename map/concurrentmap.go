package mmap

import (
	"hash/fnv"
	"sync"
)

var SEGMENT_NUM = 32

// 分段锁
type ConcurrentMap []*ConcurrentMapSegment

type ConcurrentMapSegment struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewConcurrentMap() ConcurrentMap {
	cmap := make(ConcurrentMap, 0)
	for i := 0; i < SEGMENT_NUM; i++ {
		cmap[i] = &ConcurrentMapSegment{
			data: make(map[string]interface{}),
		}
	}
	return cmap
}

func (cmap ConcurrentMap) GetSegment(key string) *ConcurrentMapSegment {
	hasher := fnv.New32()
	hasher.Write([]byte(key))
	return cmap[hasher.Sum32()%uint32(SEGMENT_NUM)]
}

func (cmap ConcurrentMap) IsEmpty() bool {
	return cmap.Size() == 0
}

func (cmap ConcurrentMap) Size() int {
	size := 0
	for i := 0; i < SEGMENT_NUM; i++ {
		segment := cmap[i]
		segment.mu.RLock()
		size += len(segment.data)
		segment.mu.RUnlock()
	}
	return size
}

func (cmap ConcurrentMap) Contain(key string) bool {
	segment := cmap.GetSegment(key)
	segment.mu.RLock()
	defer segment.mu.Unlock()
	_, ok := segment.data[key]
	return ok
}

func (cmap ConcurrentMap) Get(key string) (interface{}, bool) {
	segment := cmap.GetSegment(key)
	segment.mu.RLock()
	defer segment.mu.RUnlock()
	v, ok := segment.data[key]
	return v, ok
}

func (cmap *ConcurrentMap) Set(key string, value interface{}) {
	segment := cmap.GetSegment(key)
	segment.mu.Lock()
	defer segment.mu.Unlock()
	segment.data[key] = value
}

func (cmap *ConcurrentMap) Update(key string, value interface{}) bool {
	segment := cmap.GetSegment(key)
	segment.mu.Lock()
	defer segment.mu.Unlock()
	if _, ok := segment.data[key]; ok {
		segment.data[key] = value
		return true
	}
	return false
}

func (cmap *ConcurrentMap) Remove(key string) {
	segment := cmap.GetSegment(key)
	segment.mu.Lock()
	defer segment.mu.Unlock()
	delete(segment.data, key)
}
