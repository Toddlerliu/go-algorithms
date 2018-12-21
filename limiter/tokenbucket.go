package limiter

import (
	"sync"
	"time"
	"fmt"
)

// 漏桶算法：固定流出速率（water = max(0, water - (now - timeStamp) * rate)）
// 令牌桶算法：Token以固定的速率往桶里填充，直到达到桶的容量，多余的令牌将会被丢弃
// tokens = min(capacity, tokens + (now - timeStamp) * rate)
// 区别：漏桶算法只允许按照一定的速率发送，而令牌桶根据令牌数目来发送，最大可允许发送桶大小的流量。
type TokenBucket struct {
	sync.Mutex
	limit  float64   // 生产token速率，每秒limit个
	cap    int       // 容量
	tokens float64   // 令牌数量
	last   time.Time // TokenBucket上次更新时间
}

func NewTokenBucket(limit float64, cap int) *TokenBucket {
	return &TokenBucket{
		limit: limit,
		cap:   cap,
		last:  time.Now(),
	}
}

func (limiter *TokenBucket) Allow(n int) bool {
	limiter.Lock()
	defer limiter.Unlock()

	if limiter.cap < n {
		return false
	}
	now := time.Now()
	tokens := limiter.produce(now)
	tokens -= float64(n)
	if tokens < 0 {
		return false
	}
	limiter.last = now
	limiter.tokens = tokens
	return true
}

func (limiter *TokenBucket) AllowWithWait(n int, deadline time.Duration) (err error) {
	limiter.Lock()
	defer limiter.Unlock()

	if n > limiter.cap {
		return fmt.Errorf("input %d > limiter.cap(%d)", n, limiter.cap)
	}
	now := time.Now()
	if limiter.tokens > float64(n) {
		tokens := limiter.produce(now)
		tokens -= float64(n)
		limiter.last = time.Now()
		limiter.tokens -= tokens
		return nil
	} else {
		requirdToken := float64(n) - limiter.tokens
		requiredTime := limiter.durationFromTokens(requirdToken)
		if requiredTime > deadline {
			return fmt.Errorf("requiredTime > deadline")
		}
		t := time.NewTimer(requiredTime)
		defer t.Stop()
		select {
		case <-t.C:
			tokens := limiter.produce(now)
			tokens -= float64(n)
			limiter.last = time.Now()
			limiter.tokens -= tokens
			return nil
		}
	}
}

// 计算由时间推移导致的lim的更新状态
func (limiter *TokenBucket) produce(now time.Time) (newTokens float64) {
	last := limiter.last
	maxElapsed := limiter.durationFromTokens(float64(limiter.cap) - limiter.tokens) // 生产剩下的token需要多少时间
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}
	delta := limiter.tokensFromDuration(elapsed) // 生产tokens
	tokens := limiter.tokens + delta
	if cap := float64(limiter.cap); tokens > cap {
		tokens = cap
	}
	return tokens
}

// 计算生产tokens个需要的时间
func (limiter TokenBucket) durationFromTokens(tokens float64) time.Duration {
	seconds := tokens / float64(limiter.limit)
	return time.Nanosecond * time.Duration(1e9*seconds)
}

// 计算一段时间内能生产多少token
func (limiter TokenBucket) tokensFromDuration(d time.Duration) float64 {
	return d.Seconds() * float64(limiter.limit)
}
