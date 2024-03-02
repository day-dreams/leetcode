package hot100

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var (
		rv      *ListNode = nil
		current *ListNode = nil
	)

	for list1 != nil || list2 != nil {
		if list1 == nil {
			if rv == nil {
				return list2
			}
			current.Next = list2
			return rv
		}

		if list2 == nil {
			if rv == nil {
				return list1
			}
			current.Next = list1
			return rv
		}

		if rv == nil {
			if list1.Val > list2.Val {
				rv = &ListNode{Val: list2.Val}
				list2 = list2.Next
			} else {
				rv = &ListNode{Val: list1.Val}
				list1 = list1.Next
			}
			current = rv
			continue
		}

		if list1.Val > list2.Val {
			current.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		current = current.Next
	}

	return rv

}
