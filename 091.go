package leetcode

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		cur := s[i-1]
		if cur > '9' || cur < '0' {
			break
		}

		if cur == '0' {
			if s[i-2] > '2' || s[i-2] < '1' {
				break
			}
			dp[i] = dp[i-2]
			continue
		}

		prv := s[i-2]
		if prv == '0' {
			dp[i] = dp[i-1]
			continue
		}

		combine := s[i-2 : i]
		x, _ := strconv.Atoi(combine)
		if x <= 26 && x >= 1 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	dp[0] = 0
	fmt.Printf("dp:%+v\n", dp)
	return dp[len(s)]
}
