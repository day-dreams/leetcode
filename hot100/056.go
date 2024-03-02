package hot100

import "sort"

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] && intervals[i][1] <= intervals[j][1] {
			return true
		}
		return false
	})

	var (
		rv = [][]int{}

		current = intervals[0]
	)

	for i := 1; i < len(intervals); i++ {
		// 如果重叠，合并
		if current[1] >= intervals[i][0] {
			// 完全包裹
			if current[1] >= intervals[i][1] {
				continue
			} else {
				current[1] = intervals[i][1]
			}
		} else {
			// 否则 放入独立区间
			rv = append(rv, current)
			current = intervals[i]
		}
	}
	rv = append(rv, current)
	return rv
}
