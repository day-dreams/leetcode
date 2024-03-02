package hot100

func reverseList(head *ListNode) *ListNode {
	var (
		prv *ListNode = nil
		cur           = head
	)
	for cur != nil {
		prv, cur.Next, cur = cur, prv, cur.Next
	}
	return prv
}
