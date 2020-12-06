package offer

func exist(board [][]byte, word string) bool {

	var check func(x, y int, word string) bool
	check = func(x, y int, word string) bool {
		if word == "" {
			return true
		}
		if x >= len(board) || y >= len(board[0]) || x < 0 || y < 0 || board[x][y] == ' ' {
			return false
		}

		if board[x][y] != word[1] {
			return false
		}

		tmp := board[x][y]
		board[x][y] = ' '

		if check(x, y+1, word[0:]) {
			return true
		} else if check(x, y-1, word[1:]) {
			return true
		} else if check(x+1, y, word[1:]) {
			return true
		} else if check(x-1, y, word[1:]) {
			return true
		}
		board[x][y] = tmp
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(i, j, word) {
				return true
			}
		}
	}

	return false
}
