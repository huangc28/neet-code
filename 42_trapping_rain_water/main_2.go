package main

func trap2(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	// from left to right
	currLeftMax := 0
	maxLeft[0] = currLeftMax
	for i := 1; i < len(height); i++ {
		currLeftMax = max(currLeftMax, height[i-1])
		maxLeft[i] = currLeftMax
	}

	// from right to left
	currRightMax := 0
	maxRight[0] = currRightMax
	for j := len(height) - 2; j >= 0; j-- {
		currRightMax = max(currRightMax, height[j+1])
		maxRight[j] = currRightMax
	}

	// Calculate the total area
	totalArea := 0
	for k := 1; k < len(height)-1; k++ {
		area := min(maxLeft[k], maxRight[k]) - height[k]
		if area > 0 {
			totalArea += area
		}
	}
	return totalArea
}
