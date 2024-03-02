package leetcode

var memory = make(map[int]int)

func numTrees96(n int) int {
	if n <= 1 {
		return 1
	}
	if count, ok := memory[n]; ok {
		return count
	}
	count := 0
	for i := 1; i <= n; i++ {
		// root.val = i
		leftNum := numTrees96(i - 1)
		rightNum := numTrees96(n - i)
		count += leftNum * rightNum
	}
	memory[n] = count
	return count
}
