package main

type HeightInfo struct {
	Height int
	Index  int
}

func largestRectangleArea(heights []int) int {
	stack := make([]HeightInfo, 0, len(heights))

	largestArea := 0

	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 {
			stack = append(stack, HeightInfo{
				Height: heights[i],
				Index:  i,
			})
			continue
		}

		if heights[i] >= stack[len(stack)-1].Height {
			stack = append(stack, HeightInfo{
				Height: heights[i],
				Index:  i,
			})
			continue
		}

		var topNode HeightInfo
		for len(stack) > 0 && heights[i] <= stack[len(stack)-1].Height {
			topNode = stack[len(stack)-1]

			largestArea = max(
				largestArea,
				(i-topNode.Index)*topNode.Height,
			)
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, HeightInfo{
			Height: heights[i],
			Index:  topNode.Index,
		})
	}

	for j := 0; j < len(stack); j++ {
		area := (len(heights) - stack[j].Index) * stack[j].Height
		largestArea = max(largestArea, area)
	}

	return largestArea
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
