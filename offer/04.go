package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	x, y := len(matrix), len(matrix[0])

	row, col := 0, y-1

	for row < x && col >= 0 {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			row++
			continue
		}

		if matrix[row][col] > target {
			col--
			continue
		}

	}

	return false
}
