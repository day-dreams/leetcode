package offer

import "fmt"

func dicesProbability(n int) []float64 {

	if n <= 0 {
		return nil
	}
	rv := []float64{}
	dp := [][]float64{}
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]float64, 6*n+1))
	}

	N := n
	for n := 1; n <= N; n++ {
		for s := n; s <= 6*n; s++ {

			if s == n {
				dp[n][s] = 1
				continue
			} else if n == 1 {
				dp[n][s] = 1
				continue
			}

			// 2 3
			for i := 1; i != 7 && i <= s; i++ {
				dp[n][s] += dp[n-1][s-i]
			}

		}
	}

	sum := float64(0)

	n = N
	for i := n; i <= 6*n; i++ {
		sum += dp[n][i]
	}
	fmt.Println(dp[n])
	for i := n; i <= 6*n; i++ {
		rv = append(rv, dp[n][i]/sum)
	}
	return rv
}
