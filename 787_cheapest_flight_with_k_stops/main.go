package main

/*
distance map: foreach stop, we store the shortest distance from the priority queue.


*/
import "container/heap"

type StopInfo struct {
	dstIdx int
	cost   int
	stops  int
}

type MinHeap []StopInfo

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}
func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(StopInfo))
}

/*
(dst,cost,stops)
heap: ~(0,0,1)
*/
// const MAX_INT = int(^uint(0) >> 1)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adjMap := make(map[int][]StopInfo)
	for _, flight := range flights {
		adjMap[flight[0]] = append(adjMap[flight[0]], StopInfo{
			dstIdx: flight[1],
			cost:   flight[2],
		})
	}

	// Init minHeap to track the smallest cost to stop
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// push the first stop to heap
	heap.Push(minHeap, StopInfo{dstIdx: src, cost: 0, stops: k + 1})

	visited := make(map[StopInfo]bool)

	for minHeap.Len() > 0 {
		dstInfo := heap.Pop(minHeap).(StopInfo)

		if visited[dstInfo] {
			continue
		}

		// if we've reach the destination, simply return the ans
		if dstInfo.dstIdx == dst {
			return dstInfo.cost
		}

		visited[dstInfo] = true

		// we can not stop here. this stop is not the destination, continue popping next node
		if dstInfo.stops == 0 {
			continue
		}

		// we can stop here, add the cost of this stop.

		// push adj neighbors to minheap, we 1 stop deduction
		for _, neighborStop := range adjMap[dstInfo.dstIdx] {
			neighborStop.stops = dstInfo.stops - 1
			neighborStop.cost = neighborStop.cost + dstInfo.cost
			heap.Push(minHeap, neighborStop)
		}
	}

	return -1
}
