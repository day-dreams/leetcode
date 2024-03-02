package leetcode

func isValidBST(root *TreeNode) bool {
	var doValid func(node *TreeNode, min, max *int) bool
	doValid = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}

		// min/max == nil, 代表没有界限
		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}

		// now root is ok! lets check left/right
		leftOk := doValid(node.Left, min, &node.Val)
		if !leftOk {
			return false
		}
		return doValid(node.Right, &node.Val, max)
	}

	return doValid(root, nil, nil)
}
