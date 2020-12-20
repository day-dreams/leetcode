package tree

func maxProduct(root *TreeNode) int {
	var total int64

	if root == nil {
		return 0
	}

	subTreeSum := map[*TreeNode]int64{}
	var handle func(node *TreeNode) int64

	// calculate total
	handle = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}

		sum := int64(node.Val) + handle(node.Left) + handle(node.Right)
		subTreeSum[node] = sum
		return sum
	}

	total = handle(root)

	// divide them and calculate product
	var max int64

	handle = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}

		sub := subTreeSum[node]
		remain := total - sub

		if sub*remain > max {
			max = sub * remain
		}

		handle(node.Left)
		handle(node.Right)
		return 0
	}

	handle(root)
	if max > 0x7fffffff {
		return int(max % (1000000000 + 7))
	}
	return int(max)

}
