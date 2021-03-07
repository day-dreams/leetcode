package tree

import "fmt"

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	if root.Left != nil {
		if hasPathSum(root.Left, targetSum-root.Val) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSum(root.Right, targetSum-root.Val) {
			return true
		}
	}

	return false

	//return hasPathSumWithStack(root, targetSum)
}
func hasPathSumWithStack(root *TreeNode, targetSum int) bool {
	total := 0
	stack := []*TreeNode{}
	cur := root
	var prv *TreeNode
	for len(stack) != 0 || cur != nil {

		for cur != nil {
			total += cur.Val

			if cur.Left == nil && cur.Right == nil && total == targetSum {
				return true
			}

			stack = append(stack, cur)
			prv = cur
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		if prv == cur.Right {
			fmt.Printf("now:%v\n", cur.Val)
			stack = stack[0 : len(stack)-1]
			total -= cur.Val
			prv = cur
			cur = nil
			continue
		}

		cur = cur.Right
		prv = cur
	}

	return false
}
