package hot100

// 前缀和
func subarraySum(nums []int, k int) int {
	var (
		// 以nums[i]为结尾，和为k的子数组的个数
		count = 0

		// 从0到i的和
		prefixSum = 0

		prefixSumMap = make(map[int]int)
	)

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		if prefixSum == k {
			count++
		}
		anotherPrefixSum := prefixSum - k
		anotherCount, ok := prefixSumMap[anotherPrefixSum]
		if ok {
			count += anotherCount
		}

		prefixSumMap[prefixSum]++
	}

	return count

}

// 动态规划 o(n^2)
func subarraySum1(nums []int, k int) int {
	var (
		// 以nums[i]为结尾，和为k的子数组的个数
		counts = make([]int, len(nums))
	)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[0] == k {
				counts[0] = 1
			} else {
				counts[0] = 0
			}
			continue
		}

		// iterate from i,to i-1,i-2,.....0
		count := 0
		total := 0
		for index := i; index >= 0; index-- {
			total += nums[index]
			if total == k {
				count++
				continue
			}
		}
		counts[i] = counts[i-1] + count
	}

	return counts[len(nums)-1]
}
