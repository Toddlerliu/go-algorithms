package mybinarysearch

type node struct {
	key   string //k-v都是string，key是二叉搜索树
	value string
	left  *node
	right *node
}

// 构造节点
func newNode(key, value string) *node {
	return &node{key, value, nil, nil}
}

// 左子节点 < 当前节点 < 右子节点
type BinarySearchTree struct {
	root  *node
	count int
}

// bst 构造函数
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst BinarySearchTree) Size() int {
	return bst.count
}

func (bst BinarySearchTree) IsEmpty() bool {
	return bst.count == 0 // bst.root == nil
}

// 新增操作，若key存在更新value
func (bst *BinarySearchTree) Insert(key, value string) {
	bst.root = insert(bst.root, key, value)
}

// 在当前节点node为根节点的bst中增加下一个节点；
// 新增true；更新false
// 返回插入新节点后的bst的根（新创建的节点）
func insert(node *node, key, value string) *node {
	if node == nil { // 递归到底
		return newNode(key, value)
	}

	if key == node.key { // 更新
		node.value = value
	} else if key < node.key { // 左子树
		node.left = insert(node.left, key, value)
	} else { // 右子树
		node.right = insert(node.right, key, value)
	}
	return node
}
