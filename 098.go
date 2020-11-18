package leetcode

func isValidBST(root *TreeNode) bool {

	stack := []*TreeNode{}

	for {
		if root == nil && len(stack) == 0 {
			break
		} else if root == nil && len(stack) != 0 {
			root = stack[0]
			stack = stack[1:]
			continue
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
			if root.Left.Val > root.Val {
				return false
			}
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
			if root.Right.Val < root.Val {
				return false
			}
		}
	}

	return true
}
