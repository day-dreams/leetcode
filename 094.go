package leetcode

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal2(root)
}

// 递归
func inorderTraversal1(root *TreeNode) []int {
	var rv []int
	if root == nil {
		return nil
	}
	rv = append(rv, inorderTraversal1(root.Left)...)
	rv = append(rv, root.Val)
	rv = append(rv, inorderTraversal1(root.Right)...)
	return rv
}

// 非递归
func inorderTraversal2(root *TreeNode) []int {
	var (
		rv      []int
		parents []*TreeNode
		current *TreeNode = root
	)
	if root == nil {
		return nil
	}
	for current != nil {
		parents = append(parents, current)
		current = current.Left
	}

	for len(parents) != 0 {

		// pop one
		current = parents[len(parents)-1]
		parents = parents[0 : len(parents)-1]
		rv = append(rv, current.Val)
		current = current.Right

		// continue to leftest
		for current != nil {
			parents = append(parents, current)
			current = current.Left
		}
	}

	return rv
}
