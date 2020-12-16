package leetcode

import (
	"fmt"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	var rv []float64

	distances := map[string]map[string]float64{}
	roots := map[string]string{}

	var getRoot func(aim string) (float64, string)
	getRoot = func(aim string) (float64, string) {
		if roots[aim] == aim {
			return 1, aim
		}
		x := distances[aim][roots[aim]]
		y, root := getRoot(roots[aim])
		return x * y, root
	}

	for i, equation := range equations {

		x, y := equation[0], equation[1]

		// 没见过的节点
		if roots[x] == "" {
			roots[x] = x
		}
		if roots[y] == "" {
			roots[y] = y
		}

		disX, rootX := getRoot(x)
		_, rootY := getRoot(y)

		if rootX == rootY {
			continue
		}

		if distances[x] == nil {
			distances[x] = map[string]float64{}
		}
		if distances[y] == nil {
			distances[y] = map[string]float64{}
		}

		roots[rootX] = y
		distances[rootX][y] = values[i] / disX

		fmt.Printf("roots:%+v\n", roots)
		fmt.Printf("distances:%+v\n\n", distances)
	}

	for _, query := range queries {
		x, y := query[0], query[1]
		if roots[x] == "" || roots[y] == "" {
			rv = append(rv, -1)
			continue
		}
		disX, rootX := getRoot(x)
		disY, rootY := getRoot(y)
		if rootX != rootY {
			rv = append(rv, -1)
			continue
		}

		rv = append(rv, disX/disY)
	}

	return rv
}
