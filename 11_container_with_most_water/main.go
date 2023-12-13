package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxAmount := 0

	for l < r {
		currAmount := (r - l) * min(height[l], height[r])

		maxAmount = max(maxAmount, currAmount)

		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}

	return maxAmount
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
