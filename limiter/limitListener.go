package limiter

import "sync"

// 限制上限
type LimitListener struct {
	sem       chan struct{} // 当前数量（=上限 阻塞）
	done      chan struct{} // 限流器关闭
	closeOnce sync.Once
}

func NewLimitListener(n int) *LimitListener {
	return &LimitListener{
		sem:  make(chan struct{}, n),
		done: make(chan struct{}),
	}
}

func (limiter *LimitListener) Acquire() bool {
	select {
	case <-limiter.done:
		return false
	case limiter.sem <- struct{}{}: // sem满：阻塞
		return true
	}
}

func (limiter *LimitListener) Release() {
	<-limiter.sem
}

func (limiter *LimitListener) Close() {
	limiter.closeOnce.Do(func() {
		close(limiter.done)
	})
}
