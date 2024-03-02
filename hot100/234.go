package hot100

func isPalindrome(head *ListNode) bool {
	var (
		stack = []int{}
	)
	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}

	for i := 0; i < len(stack)/2; i++ {
		begin := i
		end := len(stack) - 1 - i
		if begin == end {
			return true
		}
		if stack[begin] != stack[end] {
			return false
		}
	}
	return true
}
