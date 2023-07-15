package main

func solve(board [][]byte) {
	var dfs func(y, x int, board [][]byte)
	dfs = func(y, x int, board [][]byte) {
		if y < 0 || x >= len(board[0]) || y >= len(board) || x < 0 || board[y][x] == '*' || board[y][x] == 'X' {
			return
		}

		board[y][x] = '*'

		dfs(y-1, x, board)
		dfs(y, x+1, board)
		dfs(y+1, x, board)
		dfs(y, x-1, board)
	}

	for x := 0; x < len(board[0]); x++ {
		dfs(0, x, board)
		dfs(len(board)-1, x, board)
	}

	for y := 0; y < len(board); y++ {
		dfs(y, 0, board)
		dfs(y, len(board[0])-1, board)
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == '0' {
				board[y][x] = 'X'
			}

			if board[y][x] == '*' {
				board[y][x] = '0'
			}
		}
	}
}
