package cache

import (
	"testing"
	"time"
	"log"
)

func TestMemCache(t *testing.T) {
	memcache := NewMemCache()
	memcache.StartAndGC(2)
	memcache.Put("hello", "world", time.Duration(1*time.Second))
	log.Println(memcache.Get("hello"))
	time.Sleep(2 * time.Second)
	log.Println(memcache.Get("hello"))
}
