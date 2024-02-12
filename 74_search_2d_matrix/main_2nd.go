package main

// 2nd try:
func searchMatrix_2(matrix [][]int, target int) bool {
	targetRow := -1

	top, bottom := 0, len(matrix)-1
	for top <= bottom {
		mid := (top + bottom) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix)-1] {
			targetRow = mid
			break
		}

		if target < matrix[mid][0] {
			bottom = mid - 1

		}

		if target > matrix[mid][len(matrix[mid])-1] {
			top = mid + 1
		}
	}

	if targetRow == -1 {
		return false
	}

	left, right := 0, len(matrix[0])-1

	for left <= right {
		mid := (left + right) / 2
		if matrix[targetRow][mid] == target {
			return true
		} else if target > matrix[targetRow][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
