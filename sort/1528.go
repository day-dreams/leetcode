package sort

func restoreString(s string, indices []int) string {

	bytes := []byte(s)
	place := map[int]byte{}
	for i, index := range indices {
		place[index] = bytes[i]
	}

	for i := 0; i < len(s); i++ {
		bytes[i] = place[i]
	}

	return string(bytes)
}
