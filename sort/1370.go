package sort

import "sort"

func sortString(s string) string {
	rv := ""

	picked := map[int]bool{}

	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	s = string(bytes)

	for len(picked) < len(s) {

		// find min
		prv := uint8(0)
		for cur := 0; cur < len(s); cur += 1 {
			if picked[cur] {
				continue
			}

			if cur == 0 {
				rv += s[0:1]
				prv = s[0]
				picked[cur] = true
				continue
			}

			if s[cur] > prv {
				rv += s[cur : cur+1]
				prv = s[cur]
				picked[cur] = true
				continue
			}

		}
		// find max
		prv = uint8(255)
		for cur := len(s) - 1; cur >= 0; cur -= 1 {
			if picked[cur] {
				continue
			}

			if cur == len(s)-1 {
				picked[cur] = true
				rv += s[cur:]
				prv = s[cur]
				continue
			}

			if s[cur] < prv {
				picked[cur] = true
				rv += s[cur : cur+1]
				prv = s[cur]
				continue
			}
		}
	}

	return rv
}
