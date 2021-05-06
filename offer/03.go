package offer

func findRepeatNumber(nums []int) int {
	memory := map[int]bool{}
	for _, num := range nums {
		if memory[num] {
			return num
		}
		memory[num] = true
	}

	return 0
}
