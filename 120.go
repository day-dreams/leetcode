package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	} else if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := [][]int{}
	for i := 0; i != len(triangle); i++ {
		dp = append(dp, make([]int, i+1))
	}

	for y := len(triangle) - 1; y >= 0; y-- {
		for x := 0; x != y+1; x++ {

			if y == len(triangle)-1 {
				dp[y][x] = triangle[y][x]
				continue
			}

			left := dp[y+1][x]
			right := dp[y+1][x+1]
			if left > right {
				dp[y][x] = right + triangle[y][x]
			} else {
				dp[y][x] = left + triangle[y][x]
			}

		}
	}

	return dp[0][0]
}
