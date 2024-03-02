package hot100

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		rv, begin, end, next, prvEnd *ListNode = nil, head, nil, nil, nil
	)
	if k <= 1 {
		return head
	}

	reverse := func(head *ListNode) (newHead *ListNode) {
		// 返回新的head
		var (
			cur            = head
			next           = cur.Next
			prv  *ListNode = nil
		)

		for i := 0; i < k; i++ {
			cur.Next, cur, next, prv = prv, next, next.Next, cur
		}
		return cur
	}

	for begin != nil {
		i := 1
		for ; i < k; i++ {
			if i == 1 {
				end = begin.Next
				continue
			}
			if end == nil {
				break
			}
			end = end.Next
		}
		if i < k || end == nil {
			break
		}
		next = end.Next
		reverse(begin)
		begin.Next = next
		if rv == nil {
			rv = end
			prvEnd = begin
		} else {
			prvEnd.Next = end
			prvEnd = begin
		}
		begin = next
	}

	if rv == nil {
		return head
	}

	if prvEnd != nil {
		prvEnd.Next = begin
	}
	return rv
}
