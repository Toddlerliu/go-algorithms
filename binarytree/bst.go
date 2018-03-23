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
	return bst.root == nil // bst.count == 0
}

// 新增操作，若key存在更新value
func (bst *BinarySearchTree) Insert(key, value string) {
	bst.root = insert(bst.root, key, value)
}

// 在当前节点node为根节点的bst中增加下一个节点；
// 返回插入新节点后的bst的根（新创建的节点）
func insert(node *node, key, value string) *node {
	if node == nil { // 递归到底
		// count ++
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

func (bst *BinarySearchTree) Get(key string) string {
	return get(bst.root, key)
}

func get(node *node, key string) string {
	if node == nil {
		return "-1"
	}

	if key == node.key {
		return node.value
	} else if key < node.key {
		return get(node.left, key)
	} else {
		return get(node.right, key)
	}
}

func (bst *BinarySearchTree) MinKey() string {
	return minKeyNode(bst.root).key
}

func (bst *BinarySearchTree) MinKeyValue() string {
	return minKeyNode(bst.root).value
}

func minKeyNode(node *node) *node {
	if node == nil {
		return nil
	}

	if node.left != nil {
		return minKeyNode(node.left)
	}
	return node
}

func (bst *BinarySearchTree) MaxKey() string {
	return maxKeyNode(bst.root).key
}

func (bst *BinarySearchTree) MaxKeyValue() string {
	return maxKeyNode(bst.root).value
}

func maxKeyNode(node *node) *node {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return maxKeyNode(node.right)
	}
	return node
}
