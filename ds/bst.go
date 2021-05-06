package ds

type node struct {
	val         int
	left, right *node
}
type BinarySearchTree struct {
	root *node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Add(val int) {
	if bst.root == nil {
		bst.root = &node{
			val:   val,
			left:  nil,
			right: nil,
		}
		return
	}

	cur := bst.root
	for {
		if cur.val >= val {
			if cur.left == nil {
				cur.left = &node{
					val:   val,
					left:  nil,
					right: nil,
				}
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &node{
					val:   val,
					left:  nil,
					right: nil,
				}
				return
			}
			cur = cur.right
		}
	}

}

func (bst *BinarySearchTree) traverseIterateWay(root *node, callback func(val int)) {

	if root == nil {
		return
	}

	var (
		stack []*node //模拟递归方式里的函数栈
		cur   = root
	)
	for cur != nil || len(stack) != 0 {

		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			callback(cur.val)
			cur = cur.right
		}

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}

	}

}

func (bst *BinarySearchTree) traverse(root *node, callback func(val int)) {
	if root == nil {
		return
	}
	bst.traverse(root.left, callback)
	callback(root.val)
	bst.traverse(root.right, callback)
}

func (bst *BinarySearchTree) Traverse(callback func(val int)) {
	//bst.traverse(bst.root, callback)
	bst.traverseIterateWay(bst.root, callback)
}
