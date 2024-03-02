package hot100

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var (
		rv     = []int{}
		x, y   = 0, 0
		top    = 0
		left   = 0
		right  = len(matrix[0]) - 1
		bottom = len(matrix) - 1
	)
	nextPosition := func(x, y int) (nextX, nextY int) {
		// left to right
		if x == top && y >= left {
			if y == right {
				return x + 1, y
			}
			return x, y + 1
		}

		// right top to right bottom
		if y == right && x >= top {
			if x == bottom {
				return x, y - 1
			}
			return x + 1, y
		}

		// right bottom to left bottom
		if x == bottom && y <= right {
			if y == left {
				if x-1 == top {
					left++
					top++
					right--
					bottom--
					return x, y + 1
				}
				return x - 1, y
			}
			return x, y - 1
		}

		if x == top+1 {
			left++
			top++
			right--
			bottom--
			return x, y + 1
		}
		return x - 1, y
	}
	for x >= top && x <= bottom &&
		y >= left && y <= right &&
		top <= bottom && left <= right {
		fmt.Printf("%v,%v %v,%v %v,%v \n", x, y, top, left, bottom, right)
		rv = append(rv, matrix[x][y])
		x, y = nextPosition(x, y)
	}
	return rv
}
