package sliding_window

func maxSatisfied(customers []int, grumpy []int, X int) int {

	if len(customers) == 0 {
		return 0
	}

	cur := 0
	max := 0
	for begin, end := 0, 0; end < len(grumpy); end++ {

		if cur > max {
			max = cur
		}
		// can satisfy
		if grumpy[end] == 0 {
			cur += customers[end]
			continue
		}

		// can control
		if begin+X > end {
			cur += customers[end]
			continue
		}

		// must lose prv customers
		for grumpy[begin] == 0 {
			begin++
		}

		if begin+X > end {
			cur += customers[end]
			continue
		}

		// cannot satisfied
		cur -= customers[begin]
		cur += customers[end]
		begin++
		//fmt.Printf("%v,%v\n", begin, end)
	}

	if cur > max {
		max = cur
	}

	return max
}
