package hot100

// 动态规划
// i能接多少水，取决于两侧最高的高度是多少。 如果最高高度小于等于i，那么接不了水；否则可以
func trap(height []int) int {
	var (
		leftMostHigh  = make([]int, len(height))
		rightMostHigh = make([]int, len(height))
		total         = 0
	)
	leftMostHigh[0] = height[0]
	for i := 1; i < len(height); i++ {
		if leftMostHigh[i-1] >= height[i] {
			leftMostHigh[i] = leftMostHigh[i-1]
		} else {
			leftMostHigh[i] = height[i]
		}
	}

	rightMostHigh[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if rightMostHigh[i+1] > height[i] {
			rightMostHigh[i] = rightMostHigh[i+1]
		} else {
			rightMostHigh[i] = height[i]
		}
	}

	// fmt.Printf("height: %+v\n", leftMostHigh)
	// fmt.Printf("height: %+v\n", rightMostHigh)
	for i := 0; i < len(height); i++ {
		if leftMostHigh[i] < height[i] {
			continue
		}
		if rightMostHigh[i] < height[i] {
			continue
		}

		if leftMostHigh[i] > rightMostHigh[i] {
			total += rightMostHigh[i] - height[i]
		} else {
			total += leftMostHigh[i] - height[i]
		}
	}

	return total
}
