package leetcode

func maxSumAfterPartitioning(arr []int, k int) int {
	if k <= 0 {
		return 0
	} else if len(arr) <= 0 {
		return 0
	}

	dp := make([]int, len(arr)+1)

	for x := 1; x <= len(arr); x++ {
		if x == 1 {
			dp[x] = arr[x-1]
			continue
		}

		sum := dp[x-1] + arr[x-1]
		for i := 1; i <= k && x-i >= 0; i++ {
			// find max one in [x-1-(i-1),x-1]
			max := arr[x-1]
			for j := x - 1 - i + 1; j != x-1; j++ {
				if arr[j] > max {
					max = arr[j]
				}
			}
			if dp[x-i]+i*max > sum {
				sum = dp[x-i] + i*max
			}
		}

		dp[x] = sum
	}

	return dp[len(arr)]
}
