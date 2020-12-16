package leetcode

func regionsBySlashes(grid []string) int {

	w, h := len(grid), len(grid[0])

	count := 0
	var inject func(x, y int)
	inject = func(x, y int) {
		if x < 0 || y < 0 || x >= w || y >= h {
			return
		}

		if grid[x][y] != ' ' {
			return
		}

		s := grid[x]
		grid[x] = s[0:y] + "x" + s[y+1:]
		inject(x, y+1)
		inject(x, y-1)
		inject(x+1, y)
		inject(x-1, y)
	}

	for i := 0; i != w; i++ {
		for j := 0; j != h; j++ {
			if grid[i][j] != ' ' {
				continue
			}

			inject(i, j)
			count++
		}
	}
	return count
}
