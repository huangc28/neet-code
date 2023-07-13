package main

/*
 */

func solveNQueens(n int) [][]string {
	colMap := make(map[int]int)  // record column occupied by the previous queen. when column x is occupied by a queen, all grid in the same column can not be occupied
	posDiag := make(map[int]int) // (r + c)
	negDiag := make(map[int]int) // (r - c)
	res := make([][]string, 0)
	board := make([][]byte, n)

	for i := 0; i < len(board); i++ {
		board[i] = make([]byte, n)
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			board[y][x] = '.'
		}
	}

	var backtracking func(y int)
	backtracking = func(y int) {
		if y == n {
			sol := getSolution(board)
			res = append(res, sol)
			return
		}

		for x := 0; x < len(board[y]); x++ {
			if colMap[x] == 1 || posDiag[y+x] == 1 || negDiag[y-x] == 1 {
				continue
			}
			colMap[x], posDiag[y+x], negDiag[y-x] = 1, 1, 1
			board[y][x] = 'Q'
			backtracking(y + 1)
			colMap[x], posDiag[y+x], negDiag[y-x] = 0, 0, 0
			board[y][x] = '.'
		}
	}

	backtracking(0)

	return res
}

func getSolution(board [][]byte) []string {
	res := make([]string, 0)
	for y := 0; y < len(board); y++ {
		row := make([]byte, len(board[y]))
		for x := 0; x < len(board); x++ {
			row[x] = board[y][x]
		}
		res = append(res, string(row))
	}
	return res
}
