package tree

import "math"

func isBalanced(root *TreeNode) bool {
	return recursiveWay(root)
}

func recursiveWay(root *TreeNode) bool {
	var validate func(root *TreeNode) (int, bool)

	validate = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		left, ok := validate(root.Left)
		if !ok {
			return -1, false
		}

		right, ok := validate(root.Right)
		if !ok {
			return -1, false
		}

		max := left
		if max < right {
			max = right
		}

		if math.Abs(float64(left-right)) <= 1 {
			return max + 1, true
		}
		return max + 1, false
	}

	_, ok := validate(root)
	return ok
}
