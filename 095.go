package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	dp := make([][]*TreeNode, n+3)
	dp[0] = []*TreeNode{nil}
	dp[1] = []*TreeNode{{Val: 1}}
	dp[2] = []*TreeNode{
		{Val: 1, Right: &TreeNode{Val: 2}},
		{Val: 2, Left: &TreeNode{Val: 1}},
	}

	if n <= 2 {
		return dp[n]
	}

	for x := 3; x <= n; x++ {
		dp[x] = []*TreeNode{}
		for i := 1; i <= x; i++ {
			left, right := dp[i-1], dp[x-i]
			for _, node1 := range left {
				for _, node2 := range right {
					dp[x] = append(dp[x], &TreeNode{Val: i, Left: node1, Right: add(node2, i)})
				}
			}
		}
	}

	return dp[n]
}

func add(root *TreeNode, i int) *TreeNode {
	if root == nil {
		return nil
	}

	rv := &TreeNode{Val: root.Val + i}
	rv.Left = add(root.Left, i)
	rv.Right = add(root.Right, i)
	return rv
}
