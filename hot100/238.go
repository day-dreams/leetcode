package hot100

func productExceptSelf(nums []int) []int {
	var (
		leftNumProduct  = make([]int, len(nums))
		rightNumProduct = make([]int, len(nums))
		rv              []int
	)

	for i := 0; i < len(leftNumProduct); i++ {
		if i == 0 {
			leftNumProduct[i] = 1
			continue
		}
		leftNumProduct[i] = leftNumProduct[i-1] * nums[i-1]
	}

	for i := len(leftNumProduct) - 1; i >= 0; i-- {
		if i == len(leftNumProduct)-1 {
			leftNumProduct[i] = 1
			continue
		}
		leftNumProduct[i] = leftNumProduct[i+1] * nums[i+1]
	}

	for i := 0; i < len(leftNumProduct); i++ {
		rv = append(rv, leftNumProduct[i]*rightNumProduct[i])
	}
	return rv
}
