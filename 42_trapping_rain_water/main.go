package main

func trap(height []int) int {
	leftPointer := 0
	rightPointer := len(height) - 1

	leftMax := 0
	rightMax := 0
	totalArea := 0

	for leftPointer < rightPointer {
		if height[leftPointer] < height[rightPointer] {
			area := leftMax - height[leftPointer]
			if area > 0 {
				totalArea += area
			}
			leftMax = max(leftMax, height[leftPointer])
			leftPointer++
		} else {
			area := rightMax - height[rightPointer]
			if area > 0 {
				totalArea += area
			}
			rightMax = max(rightMax, height[rightPointer])
			rightPointer--
		}
	}

	return totalArea
}
