package hot100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		rv, cur *ListNode = nil, nil
		c                 = 0 // 进位
	)

	for l1 != nil || l2 != nil {
		if l1 == nil {
			if rv == nil {
				return l2
			}
			cur.Next = &ListNode{Val: (l2.Val + c) % 10}
			c = (l2.Val + c) / 10
			cur = cur.Next
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			if rv == nil {
				return l1
			}
			cur.Next = &ListNode{Val: (l1.Val + c) % 10}
			c = (l1.Val + c) / 10
			cur = cur.Next
			l1 = l1.Next
			continue
		}

		if cur == nil {
			rv = &ListNode{Val: (l1.Val + l2.Val + c) % 10}
			cur = rv
		} else {
			cur.Next = &ListNode{Val: (l1.Val + l2.Val + c) % 10}
			cur = cur.Next
		}
		c = (l1.Val + l2.Val + c) / 10
		l1, l2 = l1.Next, l2.Next
	}

	if rv == nil {
		return nil
	}

	if c != 0 {
		cur.Next = &ListNode{Val: c}
	}

	return rv
}
