package leetcode

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] < '1' {
				continue
			}

			count++

			var handle func(x, y int)
			handle = func(x, y int) {

				if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
					return
				}

				grid[x][y] = '0'

				handle(x, y+1)
				handle(x, y-1)
				handle(x+1, y)
				handle(x-1, y)

			}

			handle(i, j)
		}
	}

	return count
}
