package leetcode

func isValidBST(root *TreeNode) bool {

	nodes := []int{}

	var handle func(root *TreeNode)

	handle = func(root *TreeNode) {
		if root == nil {
			return
		}

		handle(root.Left)
		nodes = append(nodes, root.Val)
		handle(root.Right)
	}

	handle(root)

	if len(nodes) <= 1 {
		return true
	}

	for i, j := 0, 1; j < len(nodes); {
		if nodes[i] >= nodes[j] {
			return false
		}
		i++
		j++
	}
	return false
}
