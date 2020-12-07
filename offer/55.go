package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	handle := func() int {

		left := maxDepth(root.Left)
		right := maxDepth(root.Right)

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	handle = func() int {

		max := 0
		stack := []*TreeNode{}
		depth := 0
		for root != nil || len(stack) != 0 {
			for root != nil {
				stack = append(stack, root)
				root = root.Left
				root++
			}
			if depth > max {
				max = depth
			}

			root = stack[len(stack)-1]
			if root.Right == nil {
				stack = stack[0 : len(stack)-1]
				depth--
			} else {
				root = root.Right
			}
		}

	}

	return handle()
}
