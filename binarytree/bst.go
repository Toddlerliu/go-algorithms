package mybinarysearch

import (
	"container/list"
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
	if node.left == nil {
		return node
	}
	return minKeyNode(node.left)
}

func (bst *BinarySearchTree) MaxKey() string {
	return maxKeyNode(bst.root).key
}

func (bst *BinarySearchTree) MaxKeyValue() string {
	return maxKeyNode(bst.root).value
}

func maxKeyNode(node *node) *node {
	if node.right == nil {
		return node
	}
	return maxKeyNode(node.right)
}

// 最小值所在的节点只会有右孩子；最大值所在的节点只会有左孩子。

// 删除最小节点(key)
// 删除只有右孩子的节点
func (bst *BinarySearchTree) RemoveMin() {
	if bst.root != nil {
		bst.root = removeMin(bst.root)
	}
}

// 返回删除节点后新的二分搜索树的根
func removeMin(node *node) *node {
	if node.left == nil { // 左孩子为空，就是最小节点
		// 右节点存在，代替现在的node节点成为新的二分搜索树的根，作为原来node节点的父节点的新的左孩子
		// 右节点不存在即为nil
		rn := node.right
		return rn
	}
	node.left = removeMin(node.left)
	return node
}

// 删除最大节点(key)
// 删除只有左孩子的节点
func (bst *BinarySearchTree) RemoveMax() {
	if bst.root != nil {
		bst.root = removeMax(bst.root)
	}
}

// 返回删除节点后新的二分搜索树的根
func removeMax(node *node) *node {
	if node.right == nil { // 左孩子为空，就是最小节点
		// 右节点存在，代替现在的node节点成为新的二分搜索树的根，作为原来node节点的父节点的新的左孩子
		// 右节点不存在即为nil
		ln := node.left
		return ln
	}
	node.right = removeMax(node.right)
	return node
}

// 删除节点
func (bst *BinarySearchTree) Remove(key string) {
	remove(bst.root, key)
}

// 删除以node为根节点的二分搜索树中键值为key的节点
// 返回删除节点后新的二分搜索树的根
// O(logn)
func remove(node *node, key string) *node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	} else if key > node.key {
		node.right = remove(node.right, key)
		return node
	} else {
		// 左子节点为空（只有右孩子）、左右孩子都为空
		if node.left == nil {
			rn := node.right
			return rn
		}
		// 右孩子为空（只有左孩子）
		if node.right == nil {
			ln := node.left
			return ln
		}
		// 左右孩子节点都存在
		//delNode := node
		// d节点右子树的最小节点s 代替 删除节点d（比d大的下一个节点s）
		successor := minKeyNode(node.right)
		// 删除右子树中的最小值s，然后s代替d，s的右节点为原来d的右节点
		successor.right = removeMin(node.right)
		// 代替节点s的左节点即d的左节点
		successor.left = node.left
		return successor
	}
}

// 深度优先遍历：O(n)
// 前序遍历：先访问当前节点，再依次递归访问左右子树
func (bst *BinarySearchTree) PreOrder() []string {
	preKeys := make([]string, 0)
	preOrder(bst.root, &preKeys)
	return preKeys
}

func preOrder(node *node, preKeys *[]string) {
	if node != nil {
		*preKeys = append(*preKeys, node.key)
		preOrder(node.left, preKeys)
		preOrder(node.right, preKeys)
	}
}

// 中序遍历：先递归访问左子树，再访问自身，再递归访问右子树。即从小到大排序（左<自身<右）
func (bst *BinarySearchTree) InOrder() []string {
	inKeys := make([]string, 0)
	inOrder(bst.root, &inKeys)
	return inKeys
}

func inOrder(node *node, inKeys *[]string) {
	if node != nil {
		inOrder(node.left, inKeys)
		*inKeys = append(*inKeys, node.key)
		inOrder(node.right, inKeys)
	}
}

// 使用Node非BinarySearchTree
func (node *node) TraverseFunc(f func(*node)) {
	if node == nil{
		return
	}
	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

// 后序遍历：先递归访问左右子树，再访问自身节点。（释放节点）
func (bst *BinarySearchTree) PostOrder() []string {
	postKeys := make([]string, 0)
	postOrder(bst.root, &postKeys)
	return postKeys
}

func postOrder(node *node, postKeys *[]string) {
	if node != nil {
		postOrder(node.left, postKeys)
		postOrder(node.right, postKeys)
		*postKeys = append(*postKeys, node.key)
	}
}

// 广度优先遍历（层序）O(n)
// 利用队列
func (bst *BinarySearchTree) LevelOrder() []string {
	q := list.New()
	q.PushBack(bst.root) // 入列
	levelKeys := make([]string, 0)
	for q.Len() > 0 {
		nd := q.Front()
		if n, ok := nd.Value.(*node); ok { //出队
			levelKeys = append(levelKeys, n.key)
			q.Remove(nd)
			if n.left != nil {
				q.PushBack(n.left)
			}
			if n.right != nil {
				q.PushBack(n.right)
			}
		}
	}
	return levelKeys
}
