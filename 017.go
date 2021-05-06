package leetcode

func letterCombinations(digits string) []string {

	rv := []string{}
	if digits == "" {
		return rv
	}

	memory := ""
	var do func(i int)
	do = func(i int) {
		if i == len(digits) {
			rv = append(rv, memory)
			return
		}

		blocks := dict[digits[i:i+1]]

		for j := range blocks {

			memory += blocks[j : j+1]
			do(i + 1)
			memory = memory[0 : len(memory)-1]
		}

	}

	do(0)
	return rv

}

var dict = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
