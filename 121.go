package leetcode

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]
	for _, x := range prices {
		tmp := x - min
		if tmp > max {
			max = tmp
		}
		if min > x {
			min = x
		}
	}
	return max
}

// func maxProfit(prices []int) int {
// 	// dp[n] 前n天最大额收益
// 	dp := make([]int, len(prices))
// 	for i := 0; i < len(prices); i++ {
// 		if i == 0 {
// 			dp[i] = 0
// 			continue
// 		}
// 		min := prices[i]
// 		for j := 0; j < i; j++ {
// 			if prices[j] < min {
// 				min = prices[j]
// 			}
// 		}
// 		max := prices[i] - min
// 		if max > dp[i-1] {
// 			dp[i] = max
// 		} else {
// 			dp[i] = dp[i-1]
// 		}
// 	}
// 	return dp[len(prices)-1]
// }
