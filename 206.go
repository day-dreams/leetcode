package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prv, cur = head, head.Next
	prv.Next=nil

	for cur != nil {

		tmp := cur.Next
		cur.Next = prv
		prv = cur
		cur = tmp
	}

	return prv

}
