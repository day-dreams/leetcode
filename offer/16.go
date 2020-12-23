package offer

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	raw := n
	if n < 0 {
		n = -n
	}

	rv := float64(1)
	X := x

	for n > 0 {

		bit := n & 1
		if bit == 1 {
			rv = rv * X
		}
		X *= X

		n = n >> 1
	}

	if raw < 0 {
		return 1 / rv
	}
	return rv

}
