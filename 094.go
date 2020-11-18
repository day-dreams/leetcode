package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	rv := []int{}
	stack := []*TreeNode{}
	for {
		// push it
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		// no more elements
		if len(stack) == 0 {
			break
		}

		// pop it
		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		// print it
		rv = append(rv, root.Val)

		// continue to right
		root = root.Right
	}

	return rv
}
