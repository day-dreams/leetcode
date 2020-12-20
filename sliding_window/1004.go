package sliding_window

func longestOnes(A []int, K int) int {

	k := 0
	max := 0

	for i, j := 0, 0; j < len(A); j++ {
		if A[j] == 1 {
			sum := j - i + 1
			if sum > max {
				max = sum
			}
			continue
		}

		if k < K {
			k += 1
			sum := j - i + 1
			if sum > max {
				max = sum
			}
			continue
		}

		for i < j {
			if A[i] == 0 {
				break
			}
			i++
		}

		i += 1
		sum := j - i + 1
		if sum > max {
			max = sum
		}

	}

	return max
}
