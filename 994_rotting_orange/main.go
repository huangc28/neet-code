package main

func orangesRotting(grid [][]int) int {
	fresh := 0
	q := make([][]int, 0) // stores rotten orange positions

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 1 {
				fresh++
			}

			if grid[y][x] == 2 {
				q = append(q, []int{y, x})
			}
		}
	}

	min := 0
	for len(q) > 0 && fresh > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			node := q[0]
			q = q[1:]
			snodes := getNodes(node[0], node[1], grid)
			for _, snode := range snodes {
				sy, sx := snode[0], snode[1]
				if grid[sy][sx] == 1 {
					fresh--
					grid[sy][sx] = 2
					q = append(q, snode)
				}
			}
		}
		min++
	}

	if fresh == 0 {
		return min
	}

	return min
}

func getNodes(y, x int, grid [][]int) [][]int {
	nodes := make([][]int, 0)
	if y-1 >= 0 {
		nodes = append(nodes, []int{y - 1, x})
	}

	if x+1 < len(grid[0]) {
		nodes = append(nodes, []int{y, x + 1})
	}

	if y+1 < len(grid) {
		nodes = append(nodes, []int{y + 1, x})
	}

	if x-1 >= 0 {
		nodes = append(nodes, []int{y, x - 1})
	}
	return nodes
}
