package main

// 2nd try:
func searchMatrix_2(matrix [][]int, target int) bool {
	targetRow := -1

	for y := 0; y < len(matrix); y++ {
		if matrix[y][0] >= target && target <= matrix[y][len(matrix)-1] {
			targetRow = y

			break
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
