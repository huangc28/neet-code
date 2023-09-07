package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	dupMap := make(map[string]bool)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			val := string(board[row][col])
			if val == "." {
				continue
			}
			_, valInRow := dupMap[fmt.Sprintf("row-%d-val-%s", row, val)]
			_, valInCol := dupMap[fmt.Sprintf("col-%d-val-%s", col, val)]

			if valInRow || valInCol {
				return false
			}

			dupMap[fmt.Sprintf("row-%d-val-%s", row, val)] = true
			dupMap[fmt.Sprintf("col-%d-val-%s", col, val)] = true

			subBoardY := row / 3
			subBoardX := col / 3
			_, valInSubBoard := dupMap[fmt.Sprintf("sub-%d-%d-val-%s", subBoardY, subBoardX, val)]
			if valInSubBoard {
				return false
			}
			dupMap[fmt.Sprintf("sub-%d-%d-val-%s", subBoardY, subBoardX, val)] = true
		}
	}
	return true
}
