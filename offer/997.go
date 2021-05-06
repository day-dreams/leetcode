package offer

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {

	}
	votes := make([]int, N+1)
	for _, vote := range trust {
		votes[vote[0]] -= 1
		votes[vote[1]] += 1
	}
	for i, x := range votes {
		if x == N-1 && i != 0 {
			return i
		}
	}

	return -1
}
