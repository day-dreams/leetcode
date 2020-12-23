package sliding_window

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	max := 0
	begin := 0
	for end := 0; end < len(s); end++ {

		if end-begin > max {
			max = end - begin
		}

		if end == begin {
			continue
		}

		// find end in [begin,end)
		for i := begin; i != end; i++ {
			if s[i] == s[end] {
				begin = i + 1
				break
			}
		}

	}

	if len(s)-begin > max {
		max = len(s) - begin
	}

	return max

}
