package main

var adjs = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for y := range board {
		visited[y] = make([]bool, len(board[y]))
	}

	var cb func(i, y, x int) bool
	cb = func(i, y, x int) bool {
		if y < 0 || y >= len(board) || x < 0 || x >= len(board[0]) || visited[y][x] || word[i] != board[y][x] {
			return false
		}

		if i == len(word)-1 {
			return true
		}

		visited[y][x] = true
		// search for next character of string in board.
		for _, adj := range adjs {
			if cb(i+1, y+adj[0], x+adj[1]) {
				return true
			}
		}
		visited[y][x] = false

		return false
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if word[0] == board[y][x] && cb(0, y, x) {
				return true
			}
		}
	}

	return false
}
