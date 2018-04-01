package unionfind

// Quick Find：查找快；union操作O(n)

// 	元素 0 1 2 3 4 5 6 7 8 9
//		 -------------------
//   id	 0 1 0 1 0 1 0 1 0 1
// 每个连在一起的组有相同的id

type UnionFind1 struct {
	id    []int // id相同连接
	count int   // 元素个数
}

func NewUnionFind1(n int) *UnionFind1 {
	uf := new(UnionFind1)
	uf.count = n
	uf.id = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始化时每个元素独立一组，所有无连接
		uf.id[i] = i
	}
	return uf
}

// 参数：元素；返回：id
// O(1)
func (uf UnionFind1) Find(p int) int {
	if p >= 0 && p <= uf.count {
		return uf.id[p]
	}
	return -1
}

// 参数：元素p和q
func (uf UnionFind1) IsConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// O(n):大数据量很慢
func (uf *UnionFind1) Union(p, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)
	if pId == qId {
		return
	}
	for i := 0; i < uf.count; i++ {
		if uf.id[i] == pId {
			uf.id[i] = qId
		}
	}
}
