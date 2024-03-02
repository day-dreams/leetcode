package hot100

func genZeroRangeOnIndexIsZero(zeroBeginAt, zeroEndAt, index int) (int, int) {
	if zeroBeginAt == -1 {
		return index, index + 1
	}
	return zeroBeginAt, index + 1
}

func genZeroRangeOnIndexIsNotZero(zeroBeginAt, zeroEndAt int) (int, int) {
	if zeroBeginAt == -1 {
		return -1, -1
	}
	return zeroBeginAt + 1, zeroEndAt + 1
}

func moveZeroes(nums []int) {
	var (
		zeroBeginAt = -1
		zeroEndAt   = -1
	)
	for index := 0; index < len(nums); index++ {
		if nums[index] == 0 {
			zeroBeginAt, zeroEndAt = genZeroRangeOnIndexIsZero(zeroBeginAt, zeroEndAt, index)
		} else {
			// move current num to zero begin
			if zeroBeginAt == -1 {
				// nothing
			} else {
				nums[zeroBeginAt] = nums[index]
				nums[index] = 0
				zeroBeginAt, zeroEndAt = genZeroRangeOnIndexIsNotZero(zeroBeginAt, zeroEndAt)
			}
		}
	}
}
