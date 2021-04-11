package leetcode

func rob(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i != len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}

		if i == 1 {
			if dp[i-1] <= nums[i] {
				dp[i] = nums[i]
			} else {
				dp[i] = dp[i-1]
			}
			continue
		}

		if dp[i-1] < dp[i-2]+nums[i] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}

	}

	return dp[len(nums)-1]

}
