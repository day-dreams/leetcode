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
	if n == 1 {
		return []*TreeNode{{Val: 1}}
	}
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	return generateTree(nums)
}

func generateTree(values []int) []*TreeNode {
	if len(values) == 0 {
		return []*TreeNode{nil}
	}
	if len(values) == 1 {
		return []*TreeNode{{Val: values[0]}}
	}
	var rv []*TreeNode
	for i := 0; i < len(values); i++ {
		lefts := generateTree(values[0:i])
		rights := generateTree(values[i+1:])
		for leftIndex := 0; leftIndex < len(lefts); leftIndex++ {
			for rightIndex := 0; rightIndex < len(rights); rightIndex++ {
				root := &TreeNode{
					Val: values[i],
				}
				root.Left = lefts[leftIndex]
				root.Right = rights[rightIndex]
				rv = append(rv, root)
			}
		}
	}
	return rv
}
