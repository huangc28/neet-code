package main

func trap3(height []int) int {
	leftMaxHeights := make([]int, len(height))
	rightMaxHeights := make([]int, len(height))

	// Iterate from left, the leftMaxHeight at height[0] is going to be 0
	currLeftMaxHeight := 0
	leftMaxHeights[0] = 0
	for i := 1; i < len(height); i++ {
		currLeftMaxHeight = max(height[i-1], currLeftMaxHeight)
		leftMaxHeights[i] = currLeftMaxHeight
	}

	// Iterate from right, the rightMaxHeight at height[0] is going to be 0
	currRightMaxHeight := 0
	rightMaxHeights[len(height)-1] = 0
	for j := len(height) - 2; j >= 0; j-- {
		currRightMaxHeight := max(height[j+1], currRightMaxHeight)
		rightMaxHeights[j] = currRightMaxHeight
	}

	total := 0
	for k := 0; k < len(height)-1; k++ {
		if height[k] > leftMaxHeights[k] || height[k] > rightMaxHeights[k] {
			continue
		}
		total += min(leftMaxHeights[k], rightMaxHeights[k]) - height[k]
	}

	return total
}

func trap3_improved(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for l < r {
		// compare leftMax & rightMax
		if height[l] <= height[r] {
			// calculate water it can trap at height[l]
			currTrap := leftMax - height[l]
			if currTrap >= 0 {
				total += currTrap
			}

			// update leftMax if height[l] is larger than leftMax
			leftMax = max(leftMax, height[l])

			// move left pointer
			l += 1
		} else {
			// calculate water it can trap at height[r]
			currTrap := rightMax - height[r]
			if currTrap >= 0 {
				total += currTrap
			}

			// update rightMax if height[r] is larger than leftMax
			rightMax = max(rightMax, height[r])

			// move right pointer
			r -= 1
		}
	}

	return total
}
