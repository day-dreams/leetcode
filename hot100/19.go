package hot100

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		fast, slow           = head, head
		prv        *ListNode = nil
		step                 = 0
	)

	for ; n > 1; n-- {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	if fast == nil {
		return nil
	}

	for fast.Next != nil {
		step++
		fast = fast.Next
		if prv == nil {
			prv, slow = head, slow.Next
		} else {
			prv, slow = slow, slow.Next
		}
	}

	if step == 0 {
		return head.Next
	} else {
		prv.Next = slow.Next
	}

	return head
}
