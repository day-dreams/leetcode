package leetcode

import (
	"fmt"
)

func numTrees(n int) int {

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	if n <= 1 {
		return dp[n]
	}

	for x := 2; x <= n; x++ {
		for i := 1; i <= x; i++ {
			dp[x] += dp[x-i] * dp[i-1]
		}
	}
	fmt.Printf("%v\n", dp)
	return dp[n]
}
