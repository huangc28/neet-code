package main

import "container/heap"

type Node struct {
	Idx    int
	Weight int
}
type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}
func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}
func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(Node))
}

func networkDelayTime(times [][]int, n int, k int) int {
	// adjacency map stores (node, time_to_arrive)
	adjMap := make(map[int][]Node)
	for _, time := range times {
		adjMap[time[0]] = append(adjMap[time[0]], Node{
			Idx:    time[1],
			Weight: time[2],
		})
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// (node_idx, weight)
	heap.Push(minHeap, Node{
		Idx:    k,
		Weight: 0,
	})
	visited := make(map[int]int)
	for minHeap.Len() > 0 {
		n := heap.Pop(minHeap).(Node)

		// If this node has been visited, we ignore it
		if _, vis := visited[n.Idx]; vis {
			continue
		}

		// Set the weight to visit this node.
		visited[n.Idx] = n.Weight

		// push adjacent nodes to minHeap
		for _, adjNode := range adjMap[n.Idx] {
			heap.Push(minHeap, Node{
				Idx:    adjNode.Idx,
				Weight: adjNode.Weight + n.Weight,
			})
		}
	}

	if len(visited) < n {
		return -1
	}

	ans := 0
	for _, weight := range visited {
		if weight > ans {
			ans = weight
		}
	}
	return ans
}
