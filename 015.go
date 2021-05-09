package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	rv := [][]int{}
	for i := 0; i < len(nums); i++ {
		a := nums[i]

		if i > 0 && a == nums[i-1] {
			continue
		}

		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {

			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}

			for k > j {

				sum := a + b + nums[k]
				if sum == 0 {
					rv = append(rv, []int{a, b, nums[k]})
					k--
					break
				}

				if sum < 0 {
					break
				}

				k--
			}

		}
	}

	return rv

}
