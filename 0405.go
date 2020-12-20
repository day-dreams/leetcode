package leetcode

func isValidBST0405(root *TreeNode) bool {

	if root == nil {
		return true
	}

	values := []int{}

	cur := root
	stack := []*TreeNode{}

	for cur != nil || len(stack) != 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		values = append(values, cur.Val)
		stack = stack[0 : len(stack)-1]
		cur = cur.Right
	}

	if len(values) <= 1 {
		return true
	}

	for i := 0; i != len(values)-1; i++ {
		if values[i] >= values[i+1] {
			return false
		}
	}

	return true
}
