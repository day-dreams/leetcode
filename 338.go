package leetcode

func countBits(num int) []int {
	dp := make([]int, num+1)

	for i := 0; i <= num; i++ {
		dp[i] = -1
	}

	for i := 0; i <= num; i++ {
		if dp[i] != -1 {
			continue
		}

		if i <= 1 {
			dp[i] = i
			continue
		}

		dp[i] = dp[i/2] + 01&i

	}

	return dp[1:num]

}
