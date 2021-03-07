package tree

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var check func(x, y *TreeNode) bool
	check = func(x, y *TreeNode) bool {
		if x == nil && y != nil {
			return false
		}
		if x != nil && y == nil {
			return false
		}
		if x == nil && y == nil {
			return true
		}

		if x.Val != y.Val {
			return false
		}

		if !check(x.Left, y.Right) {
			return false
		}
		return check(x.Right, y.Left)
	}
	return check(root.Left, root.Right)

}
