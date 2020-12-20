package tree

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	width := 1
	cur := []*TreeNode{root}

	root.Val = 0
	for len(cur) != 0 {

		next := []*TreeNode{}
		for _, node := range cur {
			if node.Left != nil {
				node.Left.Val = node.Val*2 + 1
				next = append(next, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + 2
				next = append(next, node.Right)
			}
		}

		if len(next) != 0 {
			w := next[len(next)-1].Val - next[0].Val + 1
			if w > width {
				width = w
			}
		}

		cur = next
	}

	return width
}
