package limiter

import (
	"time"
	"sync"
	"container/list"
)

// 滑动窗口：时间分格，每格一个计数器
type SlidingWindowLimiter struct {
	sync.Mutex
	Cap      int // 总容量
	Segments int // 时间格
	Interval time.Duration
	ticker   *time.Ticker
	Queue    *list.List // 维护每个时间格内的数量
}

func NewSlidingWindowLimiter(cap, segs int, interval time.Duration) *SlidingWindowLimiter {
	limiter := &SlidingWindowLimiter{
		Cap:      cap,
		Segments: segs,
		Interval: interval,
		ticker:   time.NewTicker(time.Duration(int64(interval) / int64(segs))),
		Queue:    list.New(),
	}
	for i := 0; i < segs; i++ {
		limiter.Queue.PushBack(0)
	}
	go func() {
		for {
			<-limiter.ticker.C
			limiter.Lock()
			limiter.Queue.Remove(limiter.Queue.Front())
			limiter.Queue.PushBack(0)
			limiter.Unlock()
		}
	}()
	return limiter
}

func (limiter *SlidingWindowLimiter) Increase() {
	limiter.Lock()
	defer limiter.Unlock()
	limiter.Queue.Back().Value = limiter.Queue.Back().Value.(int) + 1
}

func (limiter *SlidingWindowLimiter) IsAvailable() bool {
	limiter.Lock()
	defer limiter.Unlock()
	return limiter.cur() < limiter.Cap
}

func (limiter SlidingWindowLimiter) cur() int {
	limiter.Lock()
	defer limiter.Unlock()
	sum := 0
	for e := limiter.Queue.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			sum += i
		}
	}
	return sum
}
