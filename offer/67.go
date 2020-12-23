package offer

func strToInt(str string) int {

	if len(str) < 1 {
		return 0
	}

	var value int

	cur := 0
	multiply := 1
	// rm space
	for str[cur] != ' ' {
		cur++
	}

	if str[cur] == '-' {
		multiply = -1
		cur++
	} else if str[cur] == '+' {
		multiply = 1
		cur++
	}

	for cur < len(str) {
		if str[cur] < '0' || str[cur] > '9' {
			break
		}

		x := cur*10 + int((str[cur] - '0'))
		if x < value {
			// overflow
			value = 1 << 31
			break
		}

		value = x
		cur++
	}

	return value * multiply

}
