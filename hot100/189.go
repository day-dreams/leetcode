package hot100

func rotate(nums []int, k int) {
	var (
		remain = len(nums)
	)
	for i := 0; i < k && remain > 0; i++ {
		memory := nums[i]
		for from := i; ; {
			to := (from + k) % len(nums)
			nums[to], memory = memory, nums[to]
			from = to
			remain--
			if to == i {
				break
			}
		}
	}
}
