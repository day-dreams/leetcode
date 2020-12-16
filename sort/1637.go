package sort

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	if len(points) < 2 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	max := 0
	for i := 1; i < len(points); i++ {
		value := points[i][0] - points[i-1][0]
		if value > max {
			max = value
		}
	}
	return max
}
