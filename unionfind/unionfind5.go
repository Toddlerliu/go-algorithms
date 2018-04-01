package unionfind

// Find优化：路径压缩（保持树的层数）

type UnionFind5 struct {
	parent []int // parent[i]:元素i的父亲元素
	rank   []int // rank[i]表示以i为根的集合所表示的树的层数
	count  int   // 元素个数
}

func NewUnionFind5(n int) *UnionFind5 {
	uf := new(UnionFind5)
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

// 参数：元素；返回：根节点
// O(1)
func (uf UnionFind5) Find(p int) int {
	if p >= 0 && p <= uf.count {
		for p != uf.parent[p] {
			// 压缩过程：父亲的父亲
			uf.parent[p] = uf.parent[uf.parent[p]]
			p = uf.parent[p]
		}
		return p
	}
	return -1
}

// 最优
func (uf UnionFind5) Find2(p int) int {
	if p >= 0 && p <= uf.count {
		if p != uf.parent[p] {
			// parent[i]:元素i的父亲元素
			uf.parent[p] = uf.Find2(uf.parent[p])
		}
		return uf.parent[p]
	}
	return -1
}

func (uf UnionFind5) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// 层数少的集合root→层数多的集合root，两者rank都不变
// 层数相等的两个集合，一者的rank+1
func (uf *UnionFind5) Union(p, q int) {
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
