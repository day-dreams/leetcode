package hot100

import "fmt"

func rotateMatrix(matrix [][]int) {
	var (
		matrixLevel = (len(matrix) + 1) / 2
	)

	nextPosition := func(x, y, level int) (int, int) {
		top, left := level, level
		bottom, right := len(matrix)-level-1, len(matrix)-level-1
		if x == top && y < right {
			return level + y - left, right
		}
		if y == right && x < bottom {
			return bottom, right - (x - top)
		}
		if x == bottom && y > left {
			return bottom - (right - y), left
		}
		return top, left + bottom - x
	}

	for level := 0; level < matrixLevel; level++ {
		width := len(matrix) - 2*level
		for i := 0; i < width-1; i++ {
			x, y := level, level+i
			memory := matrix[x][y]
			for {
				nextX, nextY := nextPosition(x, y, level)
				fmt.Printf("%v,%v -> %v,%v %v\n", x, y, nextX, nextY, memory)
				matrix[nextX][nextY], memory = memory, matrix[nextX][nextY]
				if nextX == level && nextY == level+i {
					break
				}
				x, y = nextX, nextY
			}
		}
	}
}
