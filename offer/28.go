package offer

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var turn func(x *TreeNode) *TreeNode
	turn = func(x *TreeNode) *TreeNode {
		if x == nil {
			return nil
		}

		rv := &TreeNode{
			Val:   x.Val,
			Left:  turn(x.Right),
			Right: turn(x.Left),
		}

		return rv
	}

	var compare func(x, y *TreeNode) bool
	compare = func(x, y *TreeNode) bool {
		if x == nil && y == nil {
			return true
		} else if x == nil {
			return false
		} else if y == nil {
			return false
		} else if x.Val != y.Val {
			return false
		} else if !compare(x.Left, y.Left) {
			return false
		}

		return compare(x.Right, y.Right)
	}

	return compare(root, turn(root))

}
