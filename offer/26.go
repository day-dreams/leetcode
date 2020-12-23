package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}

	var compare func(x, y *TreeNode) bool
	compare = func(x, y *TreeNode) bool {
		if y == nil {
			return true
		}

		if x == nil {
			return false
		}

		if x.Val != y.Val {
			return false
		}

		if !compare(x.Left, y.Left) {
			return false
		}
		return compare(x.Right, y.Right)
	}

	level := []*TreeNode{A}
	for len(level) != 0 {
		next := []*TreeNode{}
		for _, node := range level {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

			if node.Val != B.Val {
				continue
			}
			if compare(node, B) {
				return true
			}
		}
		level = next
	}
	return false

}
