package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// return mergeTwoListsStupid(l1, l2)
	return mergeTwoListsRecursive(l1, l2)
}

func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsStupid(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsStupid(l1, l2.Next)
		return l2
	}
}

func mergeTwoListsStupid(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, backup *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}

		node := &ListNode{
			Val:  0,
			Next: nil,
		}
		if l1 == nil {
			node.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				node.Val = l1.Val
				l1 = l1.Next
			} else {
				node.Val = l2.Val
				l2 = l2.Next
			}
		}

		// add node to return value
		if r == nil {
			backup = node
		} else {
			r.Next = node
		}
		r = node
	}
	return backup
}
