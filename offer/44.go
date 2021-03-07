package offer

import "fmt"

func findNthDigit(n int) int {

	cur := 0
	bits := -1

	for bits < n {

		fmt.Printf("got bits:%v\n", bits)
		str := fmt.Sprint(cur)
		if len(str)+bits < n {
			cur++
			bits += len(str)
			continue
		}

		// find x th bit of cur
		index := n - bits
		return int(str[index-1] - '0')

	}

	return 0
}
