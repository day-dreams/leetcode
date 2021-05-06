package leetcode

func maxArea(height []int) int {
	max := 0

	if len(height) < 2 {
		return 0
	}

	for i, j := 0, len(height)-1; i != j; {

		x, y := height[i], height[j]
		height := x
		if y < x {
			height = y
		}

		width := j - i
		if height*width > max {
			max = height * width
		}

		if x < y {
			i++
		} else {
			j--
		}

	}

	return max
}
