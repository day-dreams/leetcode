package leetcode

func videoStitching(clips [][]int, T int) int {
	if T <= 0 {
		return -1
	}

	MAX := T * T
	dp := make([]int, T+1)
	for i := 1; i != T+1; i++ {
		dp[i] = MAX // define MAX
	}

	for i := 1; i <= T; i++ {

		min := MAX // define 'MAX'
		for _, clip := range clips {
			if clip[0] <= i && i <= clip[1] && dp[clip[0]] != MAX {
				if dp[clip[0]]+1 < min {
					min = dp[clip[0]] + 1
				}
			}
		}

		dp[i] = min
	}

	if dp[T] == MAX {
		return -1
	}
	return dp[T]
}
