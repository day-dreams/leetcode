package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var find func(root, target *TreeNode) []*TreeNode
	find = func(root, target *TreeNode) []*TreeNode {

		if root == nil || target == nil {
			return nil
		}
		if root.Val == target.Val {
			return []*TreeNode{root}
		}

		left := find(root.Left, target)
		if len(left) != 0 {
			return append([]*TreeNode{root}, left...)
		}

		right := find(root.Right, target)
		if len(right) != 0 {
			return append([]*TreeNode{root}, right...)
		}

		return nil
	}

	pathForP, pathForQ := find(root, p), find(root, q)
	if len(pathForP) == 0 || len(pathForQ) == 0 {
		return nil
	}

	i := 0
	for ; i < len(pathForQ) && i < len(pathForP); i++ {
		if pathForQ[i] != pathForP[i] {
			return pathForQ[i-1]
		}

	}

	return pathForQ[i]
}
