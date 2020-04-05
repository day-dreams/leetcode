package leetcode

func twoSum(nums []int, target int) []int {
	// return easy(nums, target)
	return hash(nums, target)
}

func easy(nums []int, target int) []int {
	for i := 0; i != len(nums); i += 1 {
		for j := 0; j != len(nums); j += 1 {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func hash(nums []int, target int) []int {

	num2index := map[int]int{}
	for i := 0; i != len(nums); i += 1 {
		another := target - nums[i]
		if index, ok := num2index[another]; ok {
			return []int{index, i}
		}
		num2index[nums[i]] = i
	}
	return nil
}
