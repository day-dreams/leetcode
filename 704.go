package leetcode

func search(nums []int, target int) int {
	begin, end := 0, len(nums)

	for begin <= end {

		mid := (begin + end) / 2

		// 不要陷入死结
		if mid == begin {
			if target == nums[mid] {
				return mid
			}
			return -1
		}

		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			end = mid
		} else {
			begin = mid
		}

	}

	return -1
}
