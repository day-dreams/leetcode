package hot100

func twoSum(nums []int, target int) []int {
	var (
		num2index = map[int]int{}
	)
	for index, num := range nums {
		num2index[num] = index
	}
	for index, num := range nums {
		targetIndex, ok := num2index[target-num]
		if ok && targetIndex != index {
			return []int{index, targetIndex}
		}
	}
	return nil
}
