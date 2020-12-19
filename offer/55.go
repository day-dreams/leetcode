package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	var handle func(root *TreeNode) int

	handle = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		left := handle(root.Left)
		right := handle(root.Right)
		if left > right {
			return left + 1
		}
		return right + 1

	}

	handle = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		curLevel := []*TreeNode{root}
		depth := 0
		for len(curLevel) != 0 {
			nextLevel := []*TreeNode{}
			for _, node := range curLevel {
				if node.Left != nil {
					nextLevel = append(nextLevel, node.Left)
				}
				if node.Right != nil {
					nextLevel = append(nextLevel, node.Right)
				}
			}

			depth += 1
			curLevel = nextLevel
		}

		return depth
	}

	handle = func(root *TreeNode) int {

		stack := []*TreeNode{}
		cur := root
		depth := 0
		max := 0
		var prv *TreeNode

		for len(stack) != 0 || cur != nil {
			for cur != nil {
				stack = append(stack, cur)
				depth += 1
				cur = cur.Left
			}

			cur = stack[len(stack)-1]
			if cur.Right == nil || cur.Right == prv {
				// 子节点全部访问完 我们需要向上pop一层
				prv = cur
				stack = stack[:len(stack)-1]
				cur = nil

				if depth > max {
					max = depth
				}
				depth -= 1

				continue
			}

			cur = cur.Right
		}

		return max
	}

	return handle(root)

}
