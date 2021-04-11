package leetcode

import "fmt"

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount)
	for i := range dp {
		dp[i] = -1
	}
	fmt.Printf("dp:%+v\n", dp)

	for _, i := range coins {
		if i <= amount {
			dp[i-1] = 1
		}
	}
	fmt.Printf("dp:%+v\n", dp)

	for i := 1; i <= amount; i++ {

		if dp[i-1] == 1 {
			continue
		}

		min := -1
		for j := 1; j < i; j++ {

			if dp[i-j-1] == -1 {
				continue
			}
			if dp[j-1] == -1 {
				continue
			}

			x := dp[j-1] + dp[i-j-1]
			if min == -1 || min > x {
				min = x
			}

		}
		dp[i-1] = min
		fmt.Printf("dp:%v %+v\n", i, dp)
	}

	return dp[amount-1]

}
