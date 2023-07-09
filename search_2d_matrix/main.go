package main

func searchMatrix(matrix [][]int, target int) bool {
	for x := 0; x < len(matrix[0]); x++ {
		top := 0
		bottom := len(matrix) - 1

		for top <= bottom {
			mid := (bottom + top) / 2
			if matrix[mid][x] == target {
				return true
			}

			if target > matrix[mid][x] {
				top = mid + 1
			} else {
				bottom = mid - 1
			}
		}
	}

	return false
}
