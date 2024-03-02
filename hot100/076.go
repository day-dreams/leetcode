package hot100

func minWindow(s string, t string) string {
	var (
		characters                = make(map[byte]int)
		begin                     = 0
		end                       = 0
		rv                        = ""
		matchLen                  = 0
		windowRightStopAtEndTimes = 0
	)
	for i := range t {
		characters[t[i]]++
	}

	for begin < len(s) && end < len(s) {
		// fmt.Printf("debug. begin:%v end:%v matchlen:%v. chars:%+v. rv:%+v\n",
		// 	begin, end, matchLen, characters, rv)
		if matchLen == len(t) {
			// got it!
			if len(rv) == 0 {
				if end == len(s)-1 && windowRightStopAtEndTimes > 0 {
					rv = s[begin : end+1]
				} else {
					rv = s[begin:end]
				}
			} else {
				got := ""
				if windowRightStopAtEndTimes > 0 && end == len(s)-1 {
					got = s[begin : end+1]
				} else {
					got = s[begin:end]
				}
				if len(got) < len(rv) {
					rv = got
				}
			}

			// 滑动窗口左边界，左移动
			count, ok := characters[s[begin]]
			if ok {
				// 归还之
				characters[s[begin]]++
				if count >= 0 {
					// count<0说明这个字符重复出现，不纳入范围
					matchLen--
				}
			} else {
				// 无所谓
			}
			begin++
			continue
		}

		// end已经到头了，继续考察begin字符，尝试左移滑动窗口的左边界
		if windowRightStopAtEndTimes > 0 && end == len(s)-1 {
			count, ok := characters[s[begin]]
			if ok {
				//
				if count >= 0 {
					// 重复出现，不用管它
					matchLen--
				}
				characters[s[begin]]++
			} else {
				// do nothing.弹出了一个无关字符
			}
			begin++
			// fmt.Printf("move leftside from %v to %v. end:%v matchlen:%v\n",
			// 	begin-1, begin, end, matchLen)
			windowRightStopAtEndTimes++
			continue
		}

		// 否则考察end字符,，右移滑动窗口的右边界
		count, ok := characters[s[end]]
		if ok {
			if count > 0 {
				// 小于等于0说明这个字符多出现了几次，问题不大
				matchLen++
			}
			characters[s[end]]--
		}

		// 一个从未见过的字符，直接纳入滑动窗口
		if end == len(s)-1 {
			// 第一次遇到结尾
			windowRightStopAtEndTimes++
		} else {
			end++
		}
	}

	return rv
}
