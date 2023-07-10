package main

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))

	// Initialize visited.
	for y := 0; y < len(board); y++ {
		visited[y] = make([]bool, len(board[y]))
		for x := 0; x < len(board[y]); x++ {
			visited[y][x] = false
		}
	}

	// Iterate through board to see if there is a potential match of word
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			i := 0

			// If the first char matches, it might be a potential match of word
			if board[y][x] == word[i] {
				stack := make([]byte, 0)
				stack = append(stack, board[y][x])
				visited[y][x] = true

				var dfs func(idx, y, x int) bool
				dfs = func(idx, y, x int) bool {
					// if stack has match, return true
					if string(stack) == word {
						return true
					}

					// gather surrounding characters
					cords := getLocs(board, y, x)

					// filter out valid node by comparing with word[i]
					//  - board[y][x] not visited by the previous step
					//  - word[i] matches board[y][x]
					validCords := make([][]int, 0)
					for _, cord := range cords {
						y, x := cord[0], cord[1]
						if !visited[y][x] && word[idx] == board[y][x] {
							validCords = append(validCords, []int{y, x})
						}
					}

					// try all valid coordinates to find potential matches
					for _, validCord := range validCords {
						y, x := validCord[0], validCord[1]
						stack = append(stack, board[y][x])
						visited[y][x] = true

						exists := dfs(idx+1, y, x)
						if exists {
							return true
						}

						stack = stack[:len(stack)-1]
						visited[y][x] = false
					}

					return false
				}

				hasMatch := dfs(i+1, y, x)

				if hasMatch {
					return true
				}

				stack = stack[:len(stack)-1]
				visited[y][x] = false
			}
		}
	}

	return false
}

func getLocs(board [][]byte, y, x int) [][]int {
	coords := make([][]int, 0)
	// top
	if y-1 >= 0 {
		coords = append(coords, []int{y - 1, x})
	}

	// right
	if x+1 < len(board[0]) {
		coords = append(coords, []int{y, x + 1})
	}

	// bottom
	if y+1 < len(board) {
		coords = append(coords, []int{y + 1, x})
	}

	// left
	if x-1 >= 0 {
		coords = append(coords, []int{y, x - 1})
	}

	return coords
}
