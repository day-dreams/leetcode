package hot100

func detectCycle(head *ListNode) *ListNode {
	var (
		fast, slow = head, head
		t          = 0
	)
	for {
		t++
		if fast == nil || slow == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next

		if slow == fast {
			break
		}
	}
	fast, slow = head, head
	for ; t > 0; t-- {
		fast = fast.Next
	}

	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
