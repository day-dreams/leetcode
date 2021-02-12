package offer

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	rv := [][]int{}
	route := []int{} // root -> cur
	total := 0
	var handle func(cur *TreeNode)
	handle = func(cur *TreeNode) {

		fmt.Printf("route:%+v\n", route)

		if cur == nil {
			return
		}

		total += cur.Val
		route = append(route, cur.Val)

		if cur.Left == nil && cur.Right == nil { // now we are at leaf node
			if sum == total {
				// ok we got it !
				rv = append(rv, append([]int{}, route...))
			}
		}

		handle(cur.Left)
		handle(cur.Right)
		total -= cur.Val

		if len(route) <= 1 {
			route = nil
		} else {
			route = route[0 : len(route)-1]
		}
	}

	handle(root)
	return rv
}
