package main

import "container/heap"

type GridInfo struct {
	Y      int
	X      int
	Weight int
}

type MinHeap []GridInfo

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}
func (h *MinHeap) Push(v interface{}) {
	(*h) = append(*h, v.(GridInfo))
}
func (h *MinHeap) Pop() interface{} {
	node := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return node
}

func swimInWater(grid [][]int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	visited := make([][]bool, len(grid))
	for idx := range visited {
		visited[idx] = make([]bool, len(grid[idx]))
	}

	// we start from time equals 0
	t := 0

	// we start at (0, 0)
	heap.Push(minHeap, GridInfo{
		Y:      0,
		X:      0,
		Weight: 0,
	})

	// Iterate each grid, to find minimum weight among the adj neighbors grids
	for {
		// which node is has the smallest weight to travel?
		node := heap.Pop(minHeap).(GridInfo)

		// What if we popped a node that we've already visited?
		if visited[node.Y][node.X] {
			continue
		}

		// The node to travel does not have the same height as current node.
		// The amount of time we have to weight until traveling to this node
		// is the weight of the new node.
		if node.Weight > t {
			t = node.Weight
		} else {
			// We don't need to spend any wait time to travel to next node.
			// skip
		}

		// If we reach the end of the grid, break the loop
		if len(grid)-1 == node.Y && len(grid[0])-1 == node.X {
			break
		}

		// mark the current node as visited
		visited[node.Y][node.X] = true

		neighbors := getNeiborGrids(node.Y, node.X, grid)

		// filter out visited node.
		for _, neighbor := range neighbors {
			if !visited[neighbor.Y][neighbor.X] {
				heap.Push(minHeap, neighbor)
			}
		}
	}

	return t
}

func getNeiborGrids(y, x int, grid [][]int) []GridInfo {
	ns := make([]GridInfo, 0)

	// up
	if y-1 >= 0 {
		ns = append(ns, GridInfo{
			Y:      y - 1,
			X:      x,
			Weight: grid[y-1][x],
		})
	}

	// down
	if y+1 < len(grid) {
		ns = append(ns, GridInfo{
			Y:      y + 1,
			X:      x,
			Weight: grid[y+1][x],
		})
	}

	// left
	if x-1 >= 0 {
		ns = append(ns, GridInfo{
			Y:      y,
			X:      x - 1,
			Weight: grid[y][x-1],
		})
	}

	// right
	if x+1 < len(grid[0]) {
		ns = append(ns, GridInfo{
			Y:      y,
			X:      x + 1,
			Weight: grid[y][x+1],
		})
	}

	return ns
}
