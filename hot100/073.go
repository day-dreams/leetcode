package hot100

func setZeroes(matrix [][]int) {
	var (
		rowsNeetToSet = []int{}
		colsNeetToSet = []int{}
	)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowsNeetToSet = append(rowsNeetToSet, i)
				colsNeetToSet = append(colsNeetToSet, j)
			}
		}
	}
	for col := 0; col < len(matrix[0]); col++ {
		for _, row := range rowsNeetToSet {
			matrix[row][col] = 0
		}
	}
	for row := 0; row < len(matrix); row++ {
		for _, col := range colsNeetToSet {
			matrix[row][col] = 0
		}
	}
}
