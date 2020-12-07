package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeRecursive(preorder, inorder)
}

func buildTreeRecursive(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	pos := -1
	for i, i2 := range inorder {
		if i2 == preorder[0] {
			pos = i
			break
		}
	}

	root.Left = buildTree(preorder[1:pos+1], inorder[0:pos])
	root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])

	return root
}
