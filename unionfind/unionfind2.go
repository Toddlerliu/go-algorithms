package unionfind

// Quick Union

// 	元素  0 1 2 3 4 5 6 7 8 9
//		  -------------------
// parent 0 1 2 3 4 5 6 7 8 9
// parent[i]:元素i的父亲元素

type UnionFind2 struct {
	parent []int
	count  int // 元素个数
}

func NewUnionFind2(n int) *UnionFind2 {
	uf := new(UnionFind2)
	uf.count = n
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始化时，父亲元素指向自己，两两元素互不连接
		uf.parent[i] = i
	}
	return uf
}

// 参数：元素；返回：id
// O(1)
func (uf UnionFind2) Find(p int) int {
	if p >= 0 && p <= uf.count {
		for p != uf.parent[p] {
			p = uf.parent[p]
		}
		return p
	}
	return -1
}

func (uf UnionFind2) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind2) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
}
