package offer

func replaceSpace(s string) string {
	rv := ""
	for _, c := range s {
		if c == ' ' {
			rv += "%20"
		} else {
			rv += string(c)
		}
	}
	return rv
}
