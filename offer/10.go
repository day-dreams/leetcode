package offer

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := 0, 1

	for n > 1 {
		a, b = b%1000000007, (a+b)%1000000007

		n--
	}
	return b
}
