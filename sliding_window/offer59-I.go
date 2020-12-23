package sliding_window

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}

	var rv []int
	for i := 0; i != len(nums); i++ {

		if i+k > len(nums) {
			break
		}

		max := -100000000000000
		for j := 0; j != k; j++ {
			if nums[i+j] > max {
				max = nums[i+j]
			}
		}
		rv = append(rv, max)

	}
	return rv
}
