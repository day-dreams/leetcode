package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	var rv [][]int

	targets := []*TreeNode{root}
	asc := true
	for len(targets) != 0 {

		next := []*TreeNode{}
		level := []int{}

		for begin := 0; begin < len(targets); begin++ {
			if targets[begin] == nil {
				continue
			}

			if asc {
				level = append(level, targets[begin].Val)
			} else {
				level = append([]int{targets[begin].Val}, level...)
			}
			next = append(next, targets[begin].Left, targets[begin].Right)
		}

		if len(level) != 0 {
			rv = append(rv, level)
		}
		asc = !asc
		targets = next
	}

	return rv
}
