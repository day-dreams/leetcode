package offer

import "fmt"

func movingCount(m int, n int, k int) int {

	nodes := [][]bool{}

	for i := 0; i != m; i++ {
		nodes = append(nodes, make([]bool, n))
	}

	for i := 0; i != m; i++ {
		for j := 0; j != n; j++ {
			sum := 0
			for _, bit := range fmt.Sprint(i) {
				//fmt.Printf("i=%v, i-'0'=%v\n", bit, bit-'0')
				sum += int(bit - '0')
				if sum > k {
					sum = -1
					break
				}
			}
			if sum == -1 {
				break
			}

			for _, bit := range fmt.Sprint(j) {
				sum += int(bit - '0')
				if sum > k {
					sum = -1
					break
				}
			}
			if sum != -1 {
				nodes[i][j] = true
			}
			//fmt.Printf("%v,%v = %v %v\n", i, j, sum, nodes[i][j])
		}
		//fmt.Printf("%+v\n", nodes[i])
	}

	var travel func(x, y int) int
	travel = func(x, y int) int {
		if x >= m || x < 0 || y >= n || y < 0 || !nodes[x][y] {
			return 0
		}
		nodes[x][y] = false
		ok1 := travel(x, y+1)
		ok2 := travel(x+1, y)
		return ok1 + ok2 + 1
	}
	return travel(0, 0)
}
