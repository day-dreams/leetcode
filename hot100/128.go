package hot100

func longestConsecutive(nums []int) int {
	var (
		numMap      = map[int]bool{} // 每个数字是否已经进入group
		maxGroupLen = 0
	)
	for _, num := range nums {
		numMap[num] = false
	}

	for _, num := range nums {
		if numMap[num] {
			// 已经加入了group，忽略
			continue
		}

		// 把num视为中间元素，直接枚举两侧
		groupLen := 1
		numMap[num] = true
		// 右侧元素
		for target := num + 1; ; target++ {
			_, ok := numMap[target]
			if !ok {
				break
			}
			numMap[target] = true
			groupLen++
		}
		// 左侧元素
		for target := num - 1; ; target-- {
			_, ok := numMap[target]
			if !ok {
				break
			}
			numMap[target] = true
			groupLen++
		}

		if groupLen > maxGroupLen {
			maxGroupLen = groupLen
		}
	}

	return maxGroupLen
}
