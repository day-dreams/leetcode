package leetcode

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	rv := 0

	for i := 0; i != len(nums); i++ {
		if i == 0 {
			dp[i] = 1
			rv = 1
			continue
		}

		longest := 1 // {nums[i]}
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] && longest < dp[j]+1 {
				longest = dp[j] + 1
			}
		}
		dp[i] = longest

		if dp[i] > rv {
			rv = dp[i]
		}

	}
	return rv
}
