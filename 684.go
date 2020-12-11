package leetcode

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	root := []int{0}
	for i := 0; i != n; i++ {
		root = append(root, i+1)
	}

	var findRoot func(aim int) int
	findRoot = func(aim int) int {
		if root[aim] == aim {
			return aim
		}
		return findRoot(root[aim])
	}

	redunant := []int{}
	for _, edge := range edges {
		x, y := edge[0], edge[1]

		rootX, rootY := findRoot(x), findRoot(y)
		if rootX == rootY {
			redunant = []int{x, y}
			continue
		}

		root[rootX] = y
	}

	return redunant

}
