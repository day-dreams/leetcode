package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	asc := true
	level := []*TreeNode{root}
	rv := [][]int{}
	for len(level) != 0 {

		cur := []int{}
		next := []*TreeNode{}
		for _, node := range level {

			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

			if asc {
				cur = append(cur, node.Val)
			} else {
				cur = append([]int{node.Val}, cur...)
			}

		}

		asc = !asc
		rv = append(rv, cur)
		level = next

	}
	return rv
}
