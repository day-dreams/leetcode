package tree

func postorderTraversal(root *TreeNode) []int {

	rv := []int{}
	var handle func(root *TreeNode)
	handle = func(root *TreeNode) {
		if root == nil {
			return
		}

		handle(root.Left)
		handle(root.Right)
		rv = append(rv, root.Val)
	}

	handle = func(root *TreeNode) {
		if root == nil {
			return
		}

		stack := []*TreeNode{}
		cur := root
		for len(stack) != 0 || cur != nil {

			for cur != nil {
				rv = append([]int{cur.Val}, rv...)
				if cur.Left != nil {
					stack = append(stack, cur.Left)
				}
				cur = cur.Right
			}

			if len(stack) == 0 {
				return
			}
			cur = stack[len(stack)-1]
			if len(stack) == 1 {
				stack = []*TreeNode{}
			} else {
				stack = stack[0 : len(stack)-1]
			}
			rv = append([]int{cur.Val}, rv...)
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			cur = cur.Right
		}
	}

	handle = func(root *TreeNode) {

		var prv *TreeNode
		cur := root
		stack := []*TreeNode{}

		for len(stack) != 0 || cur != nil {

			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}

			cur = stack[len(stack)-1]
			if cur.Right == nil || cur.Right == prv {
				rv = append(rv, cur.Val)
				prv = cur
				cur = nil
				stack = stack[:len(stack)-1]
				continue
			}

			cur = cur.Right
		}

	}
	handle(root)
	return rv
}
