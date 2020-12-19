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

	maxDepth := 0
	depth := 0
	stack := []*TreeNode{}
	for len(stack) != 0 || root != nil {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
			depth += 1
		}

		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			if root.Left == nil && root.Right == nil && depth > maxDepth {
				maxDepth = depth
			}

			root = root.Right
			if root == nil {
				depth--
			}

		}
	}

	return maxDepth

}
