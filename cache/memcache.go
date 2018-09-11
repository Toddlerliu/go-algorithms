package cache

import (
	"time"
	"sync"
	"errors"
)

var DefaultInterval = 60

type Cache interface {
	IsExist(key string) bool
	Get(key string) interface{}
	Put(key string, val interface{}, timeout time.Duration) error
	Remove(key string) error
	ClearAll() error
	StartAndGC(interval int) error
}

type Item struct {
	value      interface{}
	createTime time.Time     // 初始时间
	lifeTime   time.Duration // 存活时间
}

// learn form beego
type MemCache struct {
	sync.RWMutex
	items    map[string]*Item
	interval int // 过期检查间隔时间(s)
	duration time.Duration
}

func NewMemCache() *MemCache {
	return &MemCache{
		items: make(map[string]*Item),
	}
}

func (e Item) isExpire() bool {
	if e.lifeTime == 0 {
		return false
	}
	return time.Now().Sub(e.createTime) > e.lifeTime
}

func (m MemCache) IsExist(key string) bool {
	m.RLock()
	defer m.RUnlock()
	if val, ok := m.items[key]; ok {
		return !val.isExpire()
	}
	return false
}

func (m *MemCache) Put(key string, value interface{}, lifeTime time.Duration) error {
	m.Lock()
	defer m.Unlock()
	item := &Item{
		value:      value,
		createTime: time.Now(),
		lifeTime:   lifeTime,
	}
	m.items[key] = item
	return nil
}

func (m *MemCache) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()
	if v, ok := m.items[key]; ok {
		return v
	}
	return nil
}

func (m *MemCache) Remove(key string) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.items[key]; !ok {
		return errors.New("not exist")
	}
	delete(m.items, key)
	if _, ok := m.items[key]; ok {
		return errors.New("delete error")
	}
	return nil
}

func (m *MemCache) ClearAll() {
	m.Lock()
	defer m.Unlock()
	m.items = make(map[string]*Item)
}

func (m *MemCache) StartAndGC(interval int) error {
	if interval == -1 {
		interval = DefaultInterval
	}
	m.interval = interval
	m.duration = time.Duration(interval) * time.Second
	go m.checkAndClearExpire()
	return nil
}

func (m *MemCache) checkAndClearExpire() {
	if m.interval < 1 {
		return
	}
	for {
		<-time.After(m.duration)
		if m.items == nil {
			return
		}
		if keys := m.expiredKeys(); len(keys) != 0 {
			m.clearItmes(keys)
		}
	}
}

func (m *MemCache) expiredKeys() (keys []string) {
	m.Lock()
	defer m.Unlock()
	for key, val := range m.items {
		if val.isExpire() {
			keys = append(keys, key)
		}
	}
	return keys
}

func (m *MemCache) clearItmes(keys []string) {
	m.Lock()
	defer m.Unlock()
	for _, key := range keys {
		delete(m.items, key)
	}
}
