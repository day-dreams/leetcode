package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	var handle func(root *TreeNode)

	handle = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left
		handle(root.Left)
		handle(root.Right)

	}

	handle(root)

	return root
}
