package leetcode

func lengthOfLongestSubstring(s string) int {
	// return lengthOfLongestSubstring1(s)
	return lengthOfLongestSubstring2(s)
}

func lengthOfLongestSubstring1(s string) int {
	// 暴力n^2

	// 以i开头的最长str有多长

	if s == "" {
		return 0
	}

	rv := 1
	for i := 0; i != len(s); i += 1 {

		max := 1
		exists := map[byte]bool{}
		exists[s[i]] = true
		for j := i + 1; j != len(s); j += 1 {
			if !exists[s[j]] {
				max += 1
				exists[s[j]] = true
			} else {
				break
			}
		}

		if max > rv {
			rv = max
		}

	}

	return rv
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring2(s string) int {
	// 滑动窗口 O(n)
	// s[i,j)是无重复的substr，那么i<j的情况不用看了，"滑向"下一个窗口

	rv, i, j := 0, 0, 0
	set := map[byte]bool{}
	for i < len(s) && j < len(s) && i <= j {

		if !set[s[j]] {
			set[s[j]] = true
			rv = max(rv, j-i+1)
			j += 1
		} else {
			delete(set, s[i])
			i += 1
		}

	}

	return rv
}
