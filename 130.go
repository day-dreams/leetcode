package leetcode

func solve(board [][]byte) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if (i != 0 && i != len(board)-1) &&
				(j != 0 && j != len(board[0])-1) {
				continue
			}

			var inject func(x, y int)
			inject = func(x, y int) {
				if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
					return
				}

				// fmt.Printf("got:%v\n", board[x][y])
				if board[x][y] != 'O' {
					return
				}

				board[x][y] = 'x'
				inject(x, y+1)
				inject(x, y-1)
				inject(x+1, y)
				inject(x-1, y)
			}

			inject(i, j)
		}
	}

	// for _, x := range board {
	// 	fmt.Printf("x:%+v\n", x)
	// }

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'x' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

}
