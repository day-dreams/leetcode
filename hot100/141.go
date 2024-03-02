package hot100

func hasCycle(head *ListNode) bool {
	var (
		fast, slow = head, head
	)

	for fast != nil && slow != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}

	return false
}
