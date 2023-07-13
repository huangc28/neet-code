package main

func numIslands(grid [][]byte) int {
	ans := 0
	visited := make([][]bool, len(grid))
	for y := 0; y < len(grid); y++ {
		visited[y] = make([]bool, len(grid[y]))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' || visited[y][x] {
				continue
			}

			var dfs func(y, x int)
			dfs = func(y, x int) {
				if y < 0 || x >= len(grid[0]) || y >= len(grid) || x < 0 {
					return
				}

				if grid[y][x] == '0' || visited[y][x] {
					return
				}

				visited[y][x] = true

				// check up, right, bottom and left
				dfs(y-1, x)
				dfs(y, x+1)
				dfs(y+1, x)
				dfs(y, x-1)
			}
			dfs(y, x)
			ans += 1
		}
	}

	return ans
}
