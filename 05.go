package leetcode

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	rv := ""

	isPalindrome := [][]bool{}
	for i := 0; i != len(s); i++ {
		isPalindrome = append(isPalindrome, make([]bool, len(s)))
	}

	for i := 0; i != len(s); i++ {

		for j := 0; i+j < len(s); j++ {

			// i是步长
			if i == 0 {
				rv = s[j : j+1]
				isPalindrome[j][j] = true
				continue
			}

			x, y := j, j+i
			if s[x] != s[y] {
				continue
			}

			if i == 1 {
				isPalindrome[x][y] = true
				if i+1 > len(rv) {
					rv = s[x : y+1]
				}
			}
			if i > 1 && isPalindrome[x+1][y-1] {
				isPalindrome[x][y] = true
				if i +1> len(rv) {
					rv = s[x : y+1]
				}
			}

		}
	}

	return rv

}
