package limiter

type Limiter interface {
	Increase()
	IsAvailable() bool
}