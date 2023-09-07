package main

import "math"

func wallsAndGates_2(rooms [][]int) {
	q := make([][]int, 0)

	// find cords of all gates
	for y := 0; y < len(rooms); y++ {
		for x := 0; x < len(rooms[y]); x++ {
			if rooms[y][x] == 0 {
				q = append(q, []int{y, x})
			}
		}
	}

	for len(q) > 0 {
		cord := q[0]
		q = q[1:]

		neighbors := getNeighbors(cord, rooms)
		for _, neighbor := range neighbors {
			if rooms[neighbor[0]][neighbor[1]] == math.MaxInt32 {
				rooms[neighbor[0]][neighbor[1]] = rooms[cord[0]][cord[1]] + 1
				q = append(q, neighbor)
			}
		}
	}
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func getNeighbors(cord []int, rooms [][]int) [][]int {
	nodes := make([][]int, 0)
	for _, direction := range directions {
		newCord := []int{
			cord[0] + direction[0],
			cord[1] + direction[1],
		}

		// if newCord out of bound we skip
		if newCord[0] < 0 || newCord[0] > len(rooms)-1 || newCord[1] < 0 || newCord[1] > len(rooms[0])-1 {
			continue
		}

		nodes = append(nodes, newCord)
	}

	return nodes
}
