package list

// O(n/2)

type skiplistNode struct {
	key      int
	value    interface{}
	backward *skiplistNode    // 前向
	level    []*skiplistLevel // 后向（跟DoublyLinkedList相比，额外的存储空间，空间换时间）
}

type skiplistLevel struct {
	forward *skiplistNode // 后继 or 后继后继后继...  跳跃
	//span    int // 计算节点排名
}
