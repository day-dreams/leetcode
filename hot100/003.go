package hot100

// 动态规划(或者叫滑动窗口？？)
// f(i) 是以s[i]结尾的，最长字串的长度，对应的字符集为m(i)
// 那么f(i+1)取决于m(i)有没有字符s(i+1)
// 	如果没有，那么f(i+1)=f(i)+1,m(i+1)={m(i),s(i)}
//	如果有，那么要从f(i)的begin开始，一直往右找到字符s(i+1)，此时再更新m(i+1)和f(i+1)
func lengthOfLongestSubstring(s string) int {
	var (
		rv                    = 0
		currentSubStrMap      = map[byte]bool{}
		longestEndWithCurrent = make([]int, len(s)) // 以substrend结尾，最长的不重复子串有多长？
		subStrBegin           = 0
	)

	for subStrEnd := 0; subStrEnd < len(s); subStrEnd++ {
		if subStrEnd == subStrBegin {
			longestEndWithCurrent[subStrEnd] = 1
			rv = 1
			currentSubStrMap = map[byte]bool{s[subStrEnd]: true}
			continue
		}

		if !currentSubStrMap[s[subStrEnd]] {
			currentSubStrMap[s[subStrEnd]] = true
			longestEndWithCurrent[subStrEnd] = longestEndWithCurrent[subStrEnd-1] + 1
			if longestEndWithCurrent[subStrEnd] > rv {
				rv = longestEndWithCurrent[subStrEnd]
			}
			continue
		}

		// 重复了，那么就移动sub str begin
		for ; subStrBegin < subStrEnd; subStrBegin++ {
			if s[subStrBegin] == s[subStrEnd] {
				break
			}
			delete(currentSubStrMap, s[subStrBegin])
		}
		subStrBegin++
		longestEndWithCurrent[subStrEnd] = subStrEnd - subStrBegin + 1
		if longestEndWithCurrent[subStrEnd] > rv {
			rv = longestEndWithCurrent[subStrEnd]
		}
	}
	return rv

}
