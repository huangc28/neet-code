package main

import (
	"fmt"
)

/*
- check each node from the first row for all nodes that can flow to pacific
- check from the first column for all nodes that can flow to pacific
- skip checking the visited node.

pacific map

	{
	 0-0: true ---> visited and can reach pacific
	 0-1: false ---> visited and can not reach pacific
	}

if coordinate does not exist in the map or the corresponding value is false,
the coordinate can not reach the pacific
*/
func pacificAtlantic(heights [][]int) [][]int {
	// remember the coordinate that can reach the pacific
	// remember the coordinate the can reach the atlantic
	pacific, atlantic := make(map[string]bool), make(map[string]bool)

	var dfs func(y, x int, visited map[string]bool, prevHeight int)
	dfs = func(y, x int, visited map[string]bool, prevHeight int) {
		// if coordinate is out of bound
		_, exists := visited[fmt.Sprintf("%d-%d", y, x)]
		if exists ||
			y < 0 ||
			x >= len(heights[0]) ||
			y >= len(heights) ||
			x < 0 ||
			prevHeight > heights[y][x] {
			return
		}

		visited[fmt.Sprintf("%d-%d", y, x)] = true

		// check north, east, south, west
		dfs(y-1, x, visited, heights[y][x])
		dfs(y, x+1, visited, heights[y][x])
		dfs(y+1, x, visited, heights[y][x])
		dfs(y, x-1, visited, heights[y][x])
	}

	// Iterate through the first row
	// check all nodes that can reach the pacific start from row 1
	// check all nodes that can reach the atlantic start from last row
	for x := 0; x < len(heights[0]); x++ {
		dfs(0, x, pacific, heights[0][x])
		dfs(len(heights)-1, x, atlantic, heights[len(heights)-1][x])
	}

	// Iterate through the first column
	for y := 0; y < len(heights); y++ {
		dfs(y, 0, pacific, heights[y][0])
		dfs(y, len(heights[0])-1, atlantic, heights[y][len(heights[0])-1])
	}

	// find union coordinates between pacific and atlantic
	res := make([][]int, 0)
	for y := 0; y < len(heights); y++ {
		for x := 0; x < len(heights[y]); x++ {
			if pacific[fmt.Sprintf("%d-%d", y, x)] && atlantic[fmt.Sprintf("%d-%d", y, x)] {
				res = append(res, []int{y, x})
			}
		}
	}

	return res
}
