package hot100

import "fmt"

// 动态规划
func maxSubArray(nums []int) int {
	var (
		rv int

		maxSum = make([]int, len(nums))
	)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxSum[i] = nums[i]
			rv = nums[i]
			continue
		}

		if maxSum[i-1]+nums[i] > nums[i] {
			maxSum[i] = maxSum[i-1] + nums[i]
		} else {
			maxSum[i] = nums[i]
		}

		if maxSum[i] > rv {
			rv = maxSum[i]
		}
	}

	return rv
}

// 前缀和
func maxSubArray1(nums []int) int {
	var (
		rv int

		prefixSum    = make([]int, len(nums))
		minPrefxiSum int
	)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			minPrefxiSum = nums[i]
			prefixSum[i] = nums[i]
			rv = nums[i]
			continue
		}
		prefixSum[i] = prefixSum[i-1] + nums[i]

		if prefixSum[i]-minPrefxiSum > rv {
			fmt.Printf("0 rv:%v, prefixsum[i]:%v\n", rv, prefixSum[i])
			rv = prefixSum[i] - minPrefxiSum
		}
		if prefixSum[i] > rv {
			fmt.Printf("1 rv:%v, prefixsum[i]:%v\n", rv, prefixSum[i])
			rv = prefixSum[i]
		}

		if prefixSum[i] < minPrefxiSum {
			minPrefxiSum = prefixSum[i]
		}

	}

	return rv
}
