package filter

import (
	"hash"
	"hash/fnv"
	"math"
)

/**
布隆过滤器(https://en.wikipedia.org/wiki/Bloom_filter、https://llimllib.github.io/bloomfilter-tutorial/zh_CN/)
误判率p，数据规模n，二进制位个数m，哈希函数个数k
公式：p=(1-e^(-kn/m))^k
m=-(nlnp)/(ln2)^2
k=m/n*ln2
*/

type BloomFilter struct {
	// 位数
	bitSize int
	// 二进制向量，或 []byte
	bits []int64
	// 哈希函数个数
	funcs int
	// 哈希函数
	hash hash.Hash32
}

/**
n 数据规模
p 误判率, 取值范围(0, 1)
*/
func NewBloomFilter(n int, p int64) *BloomFilter {
	if n <= 0 || p <= 0 || p >= 1 {
		return nil
	}
	ln2 := math.Log(2)
	bitSize := int(-(float64(n) * math.Log(float64(p))) / (ln2 * ln2))
	return &BloomFilter{
		// m=-(nlnp)/(ln2)^2
		bitSize: bitSize,
		// k=m/n*ln2
		funcs: int(float64(bitSize) * ln2 / float64(n)),
		bits:  make([]int64, bitSize+64-1/64),
		hash:  fnv.New32a(),
	}
}

func (this *BloomFilter) Put(value string) {
	this.hash.Write([]byte(value))
	hash1 := this.hash.Sum32()
	hash2 := hash1 >> 16
	for i := 1; i <= this.funcs; i++ {
		hash := int(hash1) + i*int(hash2)
		if hash < 0 {
			hash = ^hash
		}
		// 整体的位置
		index := hash % this.bitSize
		this.set(index)
	}
}

func (this *BloomFilter) Contains(value string) bool {
	this.hash.Write([]byte(value))
	hash1 := this.hash.Sum32()
	hash2 := hash1 >> 16
	for i := 1; i <= this.funcs; i++ {
		hash := int(hash1) + i*int(hash2)
		if hash < 0 {
			hash = ^hash
		}
		// 整体的位置
		index := hash % this.bitSize
		if !this.get(index) {
			return false
		}
	}
	return true
}

func (this *BloomFilter) set(index int) {
	// slot的值(64位)
	slot := this.bits[index/64]
	// slot中的bit位 eg:0100
	bitValue := 1 << (index % 64)
	// 置1
	this.bits[index/64] = slot | bitValue
}

func (this *BloomFilter) get(index int) bool {
	slot := this.bits[index/64]
	bitValue := 1 << (index % 64)
	return slot&bitValue != 0
}
