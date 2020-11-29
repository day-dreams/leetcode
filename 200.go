package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	width, height := len(grid), len(grid[0])

	count := 0

	for x := 0; x != width; x++ {
		for y := 0; y != height; y++ {
			if grid[x][y] == '1' {
				count++
			}
			inject(grid, x, y)

		}
	}

	return count
}

func inject(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
		return
	}

	grid[x][y] = '2'
	inject(grid, x, y+1)
	inject(grid, x, y-1)
	inject(grid, x+1, y)
	inject(grid, x-1, y)
}
