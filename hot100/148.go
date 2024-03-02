package hot100

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintDebug(head *ListNode) {
	s := ""
	for head != nil {
		s += fmt.Sprintf(",%v", head.Val)
		head = head.Next
	}
	fmt.Printf(s)
}

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var count func(head *ListNode) int = func(head *ListNode) int {
		rv := 0
		for ; head != nil; head = head.Next {
			rv++
		}
		return rv
	}
	var split func(head *ListNode, count int) (left, right *ListNode) = func(head *ListNode, count int) (left *ListNode, right *ListNode) {
		if count == 1 {
			return head, nil
		}
		if count == 2 {
			right = head.Next
			left = head
			left.Next = nil
			return left, right
		}
		var tail *ListNode
		for i := 0; i < count/2; i++ {
			if left == nil {
				left = head
			}
			tail = head
			head = head.Next
		}
		right = tail.Next
		tail.Next = nil
		return left, right
	}
	var sort func(head *ListNode) *ListNode
	sort = func(head *ListNode) *ListNode {
		total := count(head)
		if total == 1 {
			return head
		}
		left, right := split(head, total)
		sortedLeft, sortedRight := sort(left), sort(right)
		var rv, tail *ListNode
		for sortedLeft != nil && sortedRight != nil {
			if sortedLeft.Val < sortedRight.Val {
				if rv == nil {
					rv, tail, sortedLeft, sortedLeft.Next =
						sortedLeft, sortedLeft, sortedLeft.Next, nil
				} else {
					tail.Next, tail, sortedLeft, sortedLeft.Next =
						sortedLeft, sortedLeft, sortedLeft.Next, nil
				}
				continue
			}
			if rv == nil {
				rv, tail, sortedRight, sortedRight.Next = sortedRight, sortedRight, sortedRight.Next, nil
			} else {
				tail.Next, tail, sortedRight, sortedRight.Next =
					sortedRight, sortedRight, sortedRight.Next, nil
			}
		}

		if sortedLeft == nil && sortedRight == nil {
			return rv
		}

		if rv == nil {
			if sortedLeft == nil {
				return sortedRight
			}
			return sortedLeft
		}

		if sortedLeft == nil {
			tail.Next = sortedRight
		} else {
			tail.Next = sortedLeft
		}
		return rv
	}

	return sort(head)
}

// tow slow!
func sortListSlow(head *ListNode) *ListNode {
	var (
		rv *ListNode
	)
	if head == nil {
		return nil
	}

	insertHandle := func(node *ListNode) {
		if rv == nil {
			rv = &ListNode{Val: node.Val}
			return
		}

		var prv *ListNode
		for begin := rv; begin != nil; begin = begin.Next {
			if begin.Val > node.Val {
				if prv == nil {
					rv = &ListNode{Val: node.Val, Next: rv}
					return
				}

				newNode := &ListNode{Val: node.Val, Next: begin}
				prv.Next = newNode
				return
			}
			prv = begin
		}

		if rv == nil {
			rv = &ListNode{Val: node.Val, Next: rv}
			return
		}

		if prv == nil {
			rv.Next = &ListNode{Val: node.Val}
		} else {
			prv.Next = &ListNode{Val: node.Val}
		}
	}

	for head != nil {
		insertHandle(head)
		head = head.Next
	}

	return rv
}
