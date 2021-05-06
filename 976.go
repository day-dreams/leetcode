package leetcode

import (
	"sort"
)

func largestPerimeter(A []int) int {
	if len(A) <= 2 {
		return 0
	}

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	// 假设c是最长边
	// a+b>c
	// sum(a,b,c) > 2c

	for i := len(A); i >= 3; i-- {

		c := A[i-1]

		for j := i - 2; j >= 1; j-- {
			if A[j]+A[j-1] > c {
				return A[j] + A[j-1] + c
			}
		}

	}

	return 0

}
