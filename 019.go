package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast, slow := head, head
	sprv := slow

	// 快慢指针，间隔n-1
	for n > 1 {
		n -= 1
		fast = fast.Next
	}

	// 快指针走到头，慢指针就是要删除的元素
	for fast.Next != nil {
		sprv = slow
		slow, fast = slow.Next, fast.Next
	}

	// 删除第一个元素
	if slow == head {
		return head.Next
	}
	sprv.Next = sprv.Next.Next
	return head
}
