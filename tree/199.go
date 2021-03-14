package tree

func rightSideView(root *TreeNode) []int {
	now := []*TreeNode{root}
	rv := []int{}

	for len(now) != 0 {

		rv = append(rv, now[len(now)-1].Val)
		next := []*TreeNode{}

		for _, node := range now {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		now = next
	}
	return rv
}
