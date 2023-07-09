package main

type Node struct {
	Index  int
	Height int
}

// Time: O(n^2)
// Space: O(n)

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]Node, 0, len(heights))
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		height := heights[i]
		if len(stack) == 0 {
			stack = append(stack, Node{i, height})
		} else {
			if stack[len(stack)-1].Height > height {
				// We can not go any further right. start calculating max area
				// until I encounter a height that is smaller than me.
				var topNode Node
				for len(stack) > 0 && stack[len(stack)-1].Height > height {
					topNode = stack[len(stack)-1]
					area := (i - topNode.Index) * topNode.Height
					maxArea = max(maxArea, area)
					stack = stack[:len(stack)-1]
				}

				stack = append(stack, Node{topNode.Index, height})
			} else {
				stack = append(stack, Node{i, height})
			}
		}
	}

	for j := 0; j < len(stack); j++ {
		node := stack[j]
		area := (len(heights) - node.Index) * node.Height
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
