package main

import "math"

func wallsAndGates(rooms [][]int) {
	for y := 0; y < len(rooms); y++ {
		for x := 0; x < len(rooms[y]); x++ {
			if rooms[y][x] == 0 {
				// perform bfs
				var bfs func(y, x, dis int)
				bfs = func(y, x, dis int) {
					// if we hit a wall
					// if we hit a gate and is not the gate itself
					if y < 0 ||
						x >= len(rooms[0]) ||
						y >= len(rooms) ||
						x < 0 ||
						rooms[y][x] == -1 ||
						(rooms[y][x] == 0 && dis > 0) {
						return
					}

					// mark node distance to gate
					if rooms[y][x] == math.MaxInt32 {
						rooms[y][x] = dis
					} else if dis <= rooms[y][x] {
						rooms[y][x] = dis
					}

					bfs(y-1, x, dis+1)
					bfs(y, x+1, dis+1)
					bfs(y+1, x, dis+1)
					bfs(y, x-1, dis+1)
				}
				bfs(y, x, 0)
			}
		}
	}
}
