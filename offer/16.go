package offer

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	rv := float64(1)
	cur := float64(1)
	old := n
	if n < 0 {
		n = -n
	}

	i := 0
	for n > 0 {

		bit := n & 1

		if i == 0 {
			cur = x
		} else {
			cur = cur * cur
		}

		if bit == 1 {
			rv *= cur
		}

		i++
		n = n >> 1

	}

	if old < 0 {
		return 1 / rv
	}

	return rv

}
