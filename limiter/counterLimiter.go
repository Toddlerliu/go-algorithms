package limiter

import (
	"sync"
	"time"
)

// 计数器法，临界问题：0:59和1:00各有100次请求（200>100）（滑动窗口1格）
type CounterLimiter struct {
	sync.Mutex
	Cap      int           // 总容量
	Cur      int           // 当前容量
	Interval time.Duration // 时间段
	ticker   *time.Ticker
}

func NewCounterLimiter(cap int, interval time.Duration) *CounterLimiter {
	limiter := &CounterLimiter{
		Cap:      cap,
		Interval: interval,
		ticker:   time.NewTicker(interval),
	}
	go func() {
		for {
			<-limiter.ticker.C
			limiter.Lock()
			limiter.Cap = 0
			limiter.Unlock()
		}
	}()
	return limiter
}

func (limiter *CounterLimiter) Increase() {
	limiter.Lock()
	defer limiter.Unlock()
	limiter.Cur++
}

func (limiter *CounterLimiter) IsAvailable() bool {
	limiter.Lock()
	defer limiter.Unlock()
	return limiter.Cur < limiter.Cap
}
