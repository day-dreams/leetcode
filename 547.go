package leetcode

func findCircleNum(M [][]int) int {
	if len(M) == 1 {
		return 1
	}

	n := len(M)
	circle := make([]int, n)

	for i := 0; i != n; i++ {
		circle[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if M[i][j] != 1 {
				continue
			}

			x, y := getParent(circle, i), getParent(circle, j)
			if x == y {
				continue
			}

			circle[x] = y
		}
		// fmt.Printf("circle:%+v\n", circle)
	}

	// fmt.Printf("circle:%+v\n", circle)
	rv := 0
	for _, x := range circle {
		if x == -1 {
			rv++
		}
	}

	return rv
}

func getParent(n []int, i int) int {
	for n[i] != -1 {
		i = n[i]
	}
	return i
}
