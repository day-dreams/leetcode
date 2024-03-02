package hot100

import (
	"sort"
)

func genGroupKey(str string) string {
	input := []byte(str)
	sort.SliceStable(input, func(i, j int) bool { return input[i] < input[j] })
	return string(input)
}

func groupAnagrams(strs []string) [][]string {
	var (
		groupKey2group = map[string][]string{}
		rv             [][]string
	)
	for _, str := range strs {
		key := genGroupKey(str)
		groupKey2group[key] = append(groupKey2group[key], str)
	}
	for _, group := range groupKey2group {
		rv = append(rv, group)
	}
	return rv
}
