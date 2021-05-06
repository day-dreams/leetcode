package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	rv := []int{}
	for head != nil {
		rv = append(rv, head.Val)
		head = head.Next
	}

	first, last := 0, len(rv)-1
	for first < last {

		rv[first], rv[last] = rv[last], rv[first]

		first++
		last--
	}
	return rv
}
