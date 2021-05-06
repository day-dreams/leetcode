package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsDivide(lists)
}

func mergeKListsDivide(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	cur := []*ListNode{}

	for len(cur) != 0 {
		next := []*ListNode{}
		for i := 0; i != len(lists); i += 1 {
			if i%2 == 1 {
				continue
			}

			if i == len(lists)-1 {
				// 最后一个孤单的list
				next = append(next, lists[i])
			} else {
				next = append(next, mergeTwoLists(lists[i], lists[i+1]))
			}
		}
		lists = next
		cur = lists
	}

	return cur[0]
}
