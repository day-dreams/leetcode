package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var handle func() []int
	handle = func() []int {
		rv := []int{}
		var travel func(root *TreeNode)
		travel = func(root *TreeNode) {
			if root == nil {
				return
			}

			travel(root.Left)
			rv = append(rv, root.Val)
			travel(root.Right)
		}
		travel(root)
		return rv
	}

	handle = func() []int {
		rv := []int{}

		stack := []*TreeNode{}
		for len(stack) != 0 || root != nil {

			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}

			if len(stack) != 0 {
				root = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				rv = append(rv, root.Val)
				root = root.Right
			}
		}

		return rv
	}

	return handle()

}
