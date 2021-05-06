package offer

import "fmt"

func translateNum(num int) int {
	if num == 0 {
		return 1
	}

	target := fmt.Sprint(num)
	dp := make([]int, len(target)+1)
	dp[0] = 1
	for i := 1; i != len(target)+1; i++ {

		if i == 1 {
			dp[1] = 1
			continue
		}

		// now we at at target[i-1]
		if target[i-2] == '1' {
			if target[i-1] <= '9' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else if target[i-2] == '2' {
			if target[i-1] <= '5' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			dp[i] = dp[i-1]
		}

	}

	return dp[len(target)]
}
