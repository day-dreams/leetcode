package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// return addTwoNumbersEasy(l1, l2)
	return addTwoNumbers2(l1, l2)
}

// 补元素，不要处理哪个空的问题
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	r := l1
	sum, bit := 0, 0
	pre1, pre2 := l1, l2 // 补全节点时，需要维护连接关系
	for {
		if l1 == nil && l2 != nil {
			l1 = &ListNode{}
			pre1.Next = l1
		} else if l2 == nil && l1 != nil {
			l2 = &ListNode{}
			pre2.Next = l2
		} else if l1 == nil && l2 == nil {
			// todo  加个新node 直接返回
			if bit != 0 {
				pre1.Next = &ListNode{
					Val:  bit,
					Next: nil,
				}
			}
			return r
		}

		sum = l1.Val + l2.Val + bit
		bit = sum / 10
		l1.Val = sum % 10
		pre1, pre2 = l1, l2
		l1, l2 = l1.Next, l2.Next
	}
}
func addTwoNumbersEasy(l1 *ListNode, l2 *ListNode) *ListNode {

	bit := 0
	var r, backup *ListNode
	for {
		cur := &ListNode{}
		if l1 != nil && l2 != nil {
			sum := l1.Val + l2.Val + bit
			bit = sum / 10
			cur.Val = sum % 10
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil && l2 != nil {
			sum := l2.Val + bit
			bit = sum / 10
			cur.Val = sum % 10
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum := l1.Val + bit
			bit = sum / 10
			cur.Val = sum % 10
			l1 = l1.Next
		} else {
			if bit != 0 {
				cur.Val = bit
				r.Next = cur
			}
			break
		}
		if r == nil {
			r = cur
			backup = cur
		} else {
			r.Next = cur
			r = cur
		}

	}
	return backup
}
