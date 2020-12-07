package offer

func minArray(numbers []int) int {
	handle := func() int {
		if len(numbers) == 0 {
			return 0
		} else if len(numbers) == 1 {
			return numbers[0]
		}

		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				return numbers[i+1]
			}
		}

		return numbers[0]
	}

	handle = func() int {
		if len(numbers) == 0 {
			return 0
		} else if len(numbers) == 1 {
			return numbers[0]
		}

		begin, end := 0, len(numbers)-1
		for begin < end {
			mid := (begin + end) / 2
			if mid == begin || mid == end {
				if numbers[begin] > numbers[end] {
					return numbers[end]
				} else {
					return numbers[begin]
				}
			}

			if numbers[mid] > numbers[end] {
				begin = mid
			} else if numbers[mid] < numbers[end] {
				end = mid
			} else {
				end--
			}
		}
		return numbers[begin]
	}

	return handle()
}
