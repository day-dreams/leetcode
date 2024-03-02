package hot100

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		cur, next           = head, head.Next
		rv        *ListNode = nil
		prv       *ListNode = nil
	)

	for cur != nil && next != nil {
		if rv == nil {
			rv = next
			cur.Next, next.Next, prv = next.Next, cur, cur
			cur = cur.Next
			if cur == nil || cur.Next == nil {
				break
			}
			next = cur.Next
			continue
		}

		prv.Next, cur.Next, next.Next = cur.Next, next.Next, cur
		prv = cur
		cur = cur.Next
		if cur == nil || cur.Next == nil {
			break
		}
		next = cur.Next
	}

	if rv == nil {
		return head
	}
	return rv
}
