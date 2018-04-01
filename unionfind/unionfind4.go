package unionfind

// 基于rank的优化：
// rank[i]表示跟节点为i的树的高度

type UnionFind4 struct {
	parent []int
	rank   []int // rank[i]表示以i为根的集合所表示的树的层数
	count  int   // 元素个数
}

func NewUnionFind4(n int) *UnionFind4 {
	uf := new(UnionFind4)
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始化时，父亲元素指向自己，两两元素互不连接
		uf.parent[i] = i
		uf.rank[i] = 1 // 初始层数为1
	}
	return uf
}

// 参数：元素；返回：id
// O(1)
func (uf UnionFind4) Find(p int) int {
	if p >= 0 && p <= uf.count {
		for p != uf.parent[p] {
			p = uf.parent[p]
		}
		return p
	}
	return -1
}

func (uf UnionFind4) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// 层数少的集合root→层数多的集合root，两者rank都不变
// 层数相等的两个集合，一者的rank+1
func (uf *UnionFind4) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[p] < uf.rank[q] {
		// 集合p层数少，p的根→集合q层数多的根，层数不变
		uf.parent[pRoot] = qRoot
	} else if uf.rank[p] > uf.rank[q] {
		uf.parent[qRoot] = pRoot
	} else {
		// 两集合层数相同
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot] += 1
	}
}
