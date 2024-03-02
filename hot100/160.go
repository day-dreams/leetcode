package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		countA, countB int
	)
	for a, b := headA, headB; a != nil || b != nil; {
		if a == b {
			return a
		}
		if a != nil {
			a = a.Next
			countA++
		}
		if b != nil {
			b = b.Next
			countB++
		}
	}

	a, b := headA, headB
	for countA > countB {
		a = a.Next
		countA--
	}
	for countA < countB {
		b = b.Next
		countB--
	}

	for a != nil && b != nil {
		if a == b {
			return a
		}
		a, b = a.Next, b.Next
	}

	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	var (
		headAMap = map[*ListNode]bool{}
	)
	for head := headA; head != nil; head = head.Next {
		headAMap[head] = true
	}
	for head := headB; head != nil; head = head.Next {
		if headAMap[head] {
			return head
		}
	}
	return nil
}
