package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	bs := func(nums []int, target int) int {
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

	pos := bs(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}
	begin, end := pos, pos
	for begin >= 0 && nums[begin] == target {
		begin--
	}
	for end < len(nums) && nums[end] == target {
		end++
	}
	return []int{begin + 1, end - 1}
}
