package hot100

// 双指针
func maxArea(height []int) int {
	var (
		rv    = 0
		left  = 0
		right = len(height) - 1
	)

	for left < right {
		minHeight := height[left]
		if height[right] < minHeight {
			minHeight = height[right]
		}
		area := minHeight * (right - left)
		if area > rv {
			rv = area
		}
		// 宽缩小的情况下，肯定是移动高较小的指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return rv
}

// 暴力枚举
func maxArea1(height []int) int {
	var (
		rv = 0
	)
	for begin := 0; begin < len(height)-1; begin++ {
		heightOfLeftest := height[begin]
		if begin > 0 && heightOfLeftest < height[begin-1] {
			// 不用看了。向右移动，还比上一个迭代里left小，面积肯定变小
			continue
		}
		for end := begin + 1; end < len(height); end++ {
			area := (end - begin) * heightOfLeftest
			if height[end] < heightOfLeftest {
				area = (end - begin) * height[end]
			}
			if area > rv {
				rv = area
			}
		}
	}
	return rv
}
