package tree

func preorderTraversal(root *TreeNode) []int {
	var rv []int

	var handle func(root *TreeNode)
	handle = func(root *TreeNode) {
		if root == nil {
			return
		}
		rv = append(rv, root.Val)
		handle(root.Left)
		handle(root.Right)
	}

	handle = func(root *TreeNode) {

		if root == nil {
			return
		}
		stack := []*TreeNode{root}

		for len(stack) != 0 {
			cur := stack[0]
			rv = append(rv, cur.Val)
			stack = stack[1:]

			if cur.Right != nil {
				stack = append([]*TreeNode{cur.Right}, stack...)
			}
			if cur.Left != nil {
				stack = append([]*TreeNode{cur.Left}, stack...)
			}

		}

	}

	handle(root)
	return rv
}
