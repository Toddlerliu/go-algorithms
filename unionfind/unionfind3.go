package unionfind

// Quick Union优化：Union的时候，比较两个集合的元素个数，个数少的追加到个数多的集合的根，尽量保持层数，并不完全准确

type UnionFind3 struct {
	parent []int
	size   []int // size[i]：以i为根的集合中元素的个数
	count  int   // 元素个数
}

func NewUnionFind3(n int) *UnionFind3 {
	uf := new(UnionFind3)
	uf.count = n
	uf.parent = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始化时，父亲元素指向自己，两两元素互不连接
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// 参数：元素；返回：id
// O(1)
func (uf UnionFind3) Find(p int) int {
	if p >= 0 && p <= uf.count {
		for p != uf.parent[p] {
			p = uf.parent[p]
		}
		return p
	}
	return -1
}

func (uf UnionFind3) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// 比较两个集合的元素个数，个数少的追加到个数多的集合的根，尽量保持层数，并不完全准确
func (uf *UnionFind3) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.size[p] < uf.size[q] {
		// 集合p元素少，p的根→集合q元素多的根，集合变大
		uf.parent[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.size[pRoot] += uf.size[qRoot]
	}
}
