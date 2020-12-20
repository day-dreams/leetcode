package offer

func cuttingRope(n int) int {

	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 1

	for i := 3; i <= n; i++ {

		max := 1

		for j := 1; j != i; j++ {
			if j*dp[i-j] > max {
				max = j * dp[i-j]
			}

			if j*(i-j) > max {
				max = j * (i - j)
			}
		}

		dp[i] = max
	}

	//fmt.Printf("%+v\n", dp)

	return dp[n]

}
