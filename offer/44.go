package offer

func findNthDigit(n int) int {

	cur := 1
	x := 0

	n = n + 1
	bits := func(x int) int {

		if x == 0 {
			return 1
		}
		rv := 0
		for x > 0 {
			rv++
			x = x / 10
		}

		return rv
	}

	for cur < n {

		more := bits(x)
		if cur+more == n {
			return x % 10
		} else if cur+more < n {
			cur += more
			x++
			continue
		}

		// 从右往左数，x的第bit位
		bit := more - (n - cur)
		for bit > 0 {
			x = x / 10
			bit--
		}
		return x % 10

	}

	return 0
}
