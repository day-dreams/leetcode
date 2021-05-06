package leetcode

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	} else if len(s) == 0 {
		return true
	}

	var translate = func(target string) []int {
		memory := map[int32]int{}
		rv := []int{}
		for _, c := range target {
			num, ok := memory[c]
			if !ok {
				num = len(memory)
				memory[c] = len(memory)
			}

			rv = append(rv, num)
		}
		return rv
	}

	x, y := translate(s), translate(t)

	for i, n := range x {
		if y[i] != n {
			return false
		}
	}

	return true
}
