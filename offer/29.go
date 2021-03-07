package offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rv := []int{}

	step := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0

	visited := [][]bool{}
	for _, ints := range matrix {
		visited = append(visited, make([]bool, len(ints)))
	}

	total := len(matrix) * len(matrix[0])
	x, y := 0, 0

	for ; total > 0; total-- {
		rv = append(rv, matrix[x][y])
		visited[x][y] = true

		nX, nY := x+step[dir][0], y+step[dir][1]
		if nX >= len(matrix) || nY >= len(matrix[0]) || nX < 0 || nY < 0 || visited[nX][nY] {
			dir = (dir + 1) % 4
			nX, nY = x+step[dir][0], y+step[dir][1]
		}

		x, y = nX, nY
	}

	return rv
}
