package redpackage

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"sync"
)

type remainPackage struct {
	remainSize  int     // 个数
	remainMoney float64 // 钱 最小单位0.01
}

type RedPackage struct {
	sync.Mutex
	totalSize  int
	totalMoney float64
	remain     *remainPackage
	rand       *rand.Rand
}

func NewRedPackage(money float64, size int) *RedPackage {
	return &RedPackage{
		totalMoney: money,
		totalSize:  size,
		remain: &remainPackage{
			remainMoney: money,
			remainSize:  size,
		},
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// 随机范围在 0.01-剩下平均值*2 之间
// 期望相同，越往后方差越大
func (p *RedPackage) GetMoney() float64 {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if p.remain.remainSize == 1 {
		p.remain.remainSize--
		return p.remain.remainMoney
	}
	min := 0.01
	max := p.remain.remainMoney / float64(p.remain.remainSize) * 2
	money, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", p.rand.Float64()*max), 64)
	if money < min {
		money = min
	}
	p.remain.remainSize--
	p.remain.remainMoney -= money
	reaminMoney, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", p.remain.remainMoney), 64)
	p.remain.remainMoney = reaminMoney
	return money
}
