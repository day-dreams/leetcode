package hot100

import (
	"sort"
)

// 双指针
// 核心是：固定i，如果i,j,k之和大于0，那么j和k的距离必须越来越近！
func threeSum(nums []int) [][]int {
	var (
		rv [][]int
	)
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		for second := first + 1; second < len(nums)-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for nums[first]+nums[second]+nums[third] > 0 && third > second {
				third--
			}

			if third == second {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				rv = append(rv, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return rv
}

// 暴力循环&&提前break
func threeSum1(nums []int) [][]int {
	var (
		rv [][]int
	)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] > 0 && nums[j] >= 0 {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				total := nums[i] + nums[j] + nums[k]
				if total > 0 {
					break
				}
				if total == 0 {
					items := []int{nums[i], nums[j], nums[k]}
					rv = append(rv, items)
					break
				}
			}
		}
	}
	return rv
}
