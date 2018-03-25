package mybinarysearch

var (
	// 遍历返回
	preKeys []string = make([]string, 0)
	inKeys []string = make([]string, 0)
	postKeys []string = make([]string, 0)
)

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

// 查找：若找到返回value和true；否则""和false
func (bst *BinarySearchTree) Search(key string) (string, bool) {
	return search(bst.root, key)
}

func search(node *node, key string) (string, bool) {
	if node == nil {
		return "not found", false
	}
	if key == node.key {
		return node.value, true
	} else if key < node.key {
		return search(node.left, key)
	} else {
		return search(node.right, key)
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

func (bst *BinarySearchTree) DeleteMinNode() *node {
	return deleteMinNode(bst.root)
}

func deleteMinNode(node *node) *node {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return deleteMinNode(node.left)
	}
	return node
}

func (bst *BinarySearchTree) DeleteMaxNode() *node {
	return deleteMaxNode(bst.root)
}

func deleteMaxNode(node *node) *node {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return deleteMaxNode(node.right)
	}
	return node
}

func (bst *BinarySearchTree) Delete(key string) (*node, bool) {
	return delete(bst.root, key)
}

func delete(node *node, key string) (*node, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.key {
		return delete(node.left, key)
	} else if key > node.key {
		return delete(node.right, key)
	} else {
		// TO DO
		// 无子节点
	}
	return node, true
}

//前序遍历：先访问当前节点，再依次递归访问左右子树
func (bst *BinarySearchTree) PreOrder() []string {
	preOrder(bst.root)
	return preKeys
}

func preOrder(node *node) {
	if node != nil {
		preKeys = append(preKeys, node.key)
		preOrder(node.left)
		preOrder(node.right)
	}
}

//中序遍历：先递归访问左子树，再访问自身，再递归访问右子树。即从小到大排序（左<自身<右）
func (bst *BinarySearchTree) InOrder() []string {
	inOrder(bst.root)
	return inKeys
}

func inOrder(node *node) {
	if node != nil {
		inOrder(node.left)
		inKeys = append(inKeys, node.key)
		inOrder(node.right)
	}
}

//后序遍历：先递归访问左右子树，再访问自身节点。（释放节点）
func (bst *BinarySearchTree) PostOrder() []string {
	postOrder(bst.root)
	return postKeys
}

func postOrder(node *node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		postKeys = append(postKeys, node.key)
	}
}