package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	prv := 1
	cur := 2
	for i := 0; i < n-2; i++ {
		tmp := prv + cur
		prv, cur = cur, tmp
	}
	return cur
}
