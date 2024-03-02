package hot100

func canJump(nums []int) bool {
	var (
		dp = make([]bool, len(nums))
	)
	for i, _ := range nums {
		if i == 0 {
			dp[i] = true
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
