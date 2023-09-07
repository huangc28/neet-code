package main

func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1
	mid := 0
	for top <= bottom {
		mid = (top + bottom) / 2
		// we've found the row where target is at.
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			break
		}

		// target is larger than the largest number of current row
		// target is at bottom rows
		if target > matrix[mid][len(matrix[mid])-1] {
			top = mid + 1
		}

		// target is less than the smallest number of current row
		// target is at top rows.
		if target < matrix[mid][0] {
			bottom = mid - 1
		}

		// target is
	}

	// we've found the target row, perform ordinary binary search.
	tr := matrix[mid]

	left, right := 0, len(tr)-1
	for left <= right {
		mid := (left + right) / 2
		if tr[mid] == target {
			return true
		}
		if target > tr[mid] {
			// search right
			left = mid + 1
		} else {
			// search left
			right = mid - 1
		}
	}
	return false
}
