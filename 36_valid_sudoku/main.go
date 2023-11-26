package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	dupMap := make(map[string]bool)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == '.' {
				continue
			}

			currVal := int(board[row][col])

			rowKey := fmt.Sprintf("row-%d-%d", row, currVal)
			colKey := fmt.Sprintf("col-%d-%d", col, currVal)

			// locate subboard
			sbRow := row / 3
			sbCol := col / 3
			sbKey := fmt.Sprintf("sb-%d-%d-%d", sbRow, sbCol, currVal)

			if dupMap[rowKey] || dupMap[colKey] || dupMap[sbKey] {
				return false
			}

			dupMap[rowKey] = true
			dupMap[colKey] = true
			dupMap[sbKey] = true
		}
	}
	return true
}
