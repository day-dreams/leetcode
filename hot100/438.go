package hot100

func findAnagrams(s string, p string) []int {
	var (
		charactorMap        = map[byte]int{}
		rv                  = []int{}
		begin               = 0
		end                 = 0
		debug        func() = func() {
			// fmt.Printf("count:%v begin:%v end:%v %+v\n", charactorCount, begin, end, charactorMap)
		}
	)

	for c := range p {
		charactorMap[p[c]]++
	}

	for end < len(s) {
		debug()
		// 检查当前窗口
		// 匹配到了！
		if end-begin == len(p) {
			rv = append(rv, begin)
			charactorMap[s[begin]]++
			begin++
			continue
		}

		count, ok := charactorMap[s[end]]
		if ok {
			// 属于p的字符
			// 如果还有剩余次数，纳入窗口
			if count > 0 {
				charactorMap[s[end]]--
				end++
				continue
			}
			// 否则归还第一个字符，缩小窗口
			charactorMap[s[begin]]++
			begin++
			continue
		}

		// end是不属于p的字符，那么窗口失去意义，直接缩到最小
		// 并把窗口移到end的后面
		for begin < end {
			charactorMap[s[begin]]++
			begin++
		}

		begin++
		end++

	}

	if end-begin == len(p) {
		rv = append(rv, begin)
	}

	return rv
}
