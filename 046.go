package leetcode

func permute(nums []int) [][]int {

	var rv [][]int

	var do func(i int)
	do = func(i int) {
		// [0,i) 已经填充完毕

		if i == len(nums) {
			dst := make([]int, len(nums))
			copy(dst, nums)
			rv = append(rv, dst)
			return
		}

		for j := i; j < len(nums); j++ {

			nums[i], nums[j] = nums[j], nums[i]
			do(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}

	}

	do(0)
	return rv
}
