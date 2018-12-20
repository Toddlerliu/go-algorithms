package limiter

// 漏桶算法：固定流出速率（water = max(0, water - (now - timeStamp) * rate)）
// 令牌桶算法：Token以固定的速率往桶里填充，直到达到桶的容量，多余的令牌将会被丢弃（tokens = min(capacity, tokens + (now - timeStamp) * rate)）
// 区别：漏桶算法只允许按照一定的速率发送，而令牌桶根据令牌数目来发送，最大可允许发送桶大小的流量。
type Limiter interface {
	Increase()
	IsAvailable() bool
}