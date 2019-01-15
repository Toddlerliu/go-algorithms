package mybinarysearch

// 树状数组
//
// 3个的：                                    s8          2^3 = 1000 = 8
// 2个的：                s4                              2^2 = 100 = 4
// 1个的：      s2                  s6			   s[i]   2^1 = 10 = 2
// 0个的： s1        s3        s5        s7               2^0 = 1 = 1
//   	    1    2    3    4    5    6    7    8
// 		  0001 0010 0011 0100 0101 0110 0111 1000  a[i]
// 0个数：  0    1    0    2    0    1    0    3
// 0001 s[1] = a[1]
// 0010 s[2] = a[1]+a[2]=s[1]+a[2]
// 0011 s[3] = a[3]
// 0100 s[4] = a[1]+a[2]+a[3]+a[4] = s[2]+s[3]+a[4]
// 0101 s[5] = a[5]
// 0110 s[6] = a[5]+a[6]
// 0111 s[7] = a[7]
// 1000 s[8] = a[1]+a[2]+a[3]+a[4]+a[5]+a[6]+a[7]+a[8] = s[4]+s[6]+s[7]+a[8]
// s[i] = a[i-2^k+1]+a[i-2^k+2]+......+a[i] (k为i的二进制中从最低位到最高位连续零的个数)
type BinaryIndexedTree struct {
	a []int // 原数组
}

func NewBinaryIndexedTree(len int) *BinaryIndexedTree {
	return &BinaryIndexedTree{
		a: make([]int, len),
	}
}

// 父节点编号=当前节点n+lowbit(n)
func (t *BinaryIndexedTree) AddOrUpdate(index, value int) {
	for index <= len(t.a) {
		t.a[index] += value
		index += lowbit(index)
	}
}

// 子节点编号=当前节点n-lowbit(n)
func (t BinaryIndexedTree) Sum(index int) (sum int) {
	sum = 0
	for index > 0 {
		sum += t.a[index]
		index -= lowbit(index)
	}
	return
}

func lowbit(x int) int {
	return x & -x
}
