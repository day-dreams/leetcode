package hot100

func searchMatrix(matrix [][]int, target int) bool {

	var (
		x, y = len(matrix) - 1, 0
	)

	for {
		if x < 0 || y >= len(matrix[0]) {
			return false
		}

		if matrix[x][y] == target {
			return true
		}

		if x == 0 && y == len(matrix[0])-1 {
			return false
		}

		if x == 0 {
			if matrix[x][y+1] > target {
				return false
			}
			x, y = x, y+1
			continue
		}

		if y == len(matrix[0])-1 {
			if matrix[x-1][y] < target {
				return false
			}
			x, y = x-1, y
			continue
		}

		if matrix[x][y] > target {
			x, y = x-1, y
		} else {
			x, y = x, y+1
		}
	}

	return false
}
