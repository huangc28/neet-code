package main

func trap2(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	// from left to right
	currLeftMax := 0
	for i := 0; i < len(height); i++ {
		if i == 0 {
			maxLeft[i] = 0
			continue
		}

		currLeftMax = max(currLeftMax, height[i-1])
		maxLeft[i] = currLeftMax
	}

	// from right to left
	currRightMax := 0
	for j := len(height) - 1; j >= 0; j-- {
		if j == len(height)-1 {
			maxRight[len(height)-1] = 0
			continue
		}

		currRightMax = max(currRightMax, height[j+1])
		maxRight[j] = currRightMax
	}

	// Calculate the total area
	totalArea := 0
	for k := 0; k < len(height); k++ {
		area := min(maxLeft[k], maxRight[k]) - height[k]
		if area > 0 {
			totalArea += area
		}
	}
	return totalArea
}
