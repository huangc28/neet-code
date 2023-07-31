package main

import (
	"container/heap"
	"sort"
)

/*
- sort the intervals by start time
- sort the query in ascending order
- push interval to min heap until q is not in the range of the interval, q >= intervals[0][0] and i < len(intervals)
- Keep popping interval from heap while q is not in range i.e. q > intervals[0][1]. The first interval in the min heap would be the answer at query[i]

*/

type Interval struct {
	interval []int
	size     int
}

type IntervalHeap []Interval

func (self IntervalHeap) Len() int {
	return len(self)
}

func (self IntervalHeap) Less(i, j int) bool {
	return self[i].size < self[j].size
}

func (self IntervalHeap) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (h *IntervalHeap) Push(obj interface{}) {
	*h = append(*h, obj.(Interval))
}

func (h *IntervalHeap) Pop() interface{} {
	// swap
	// chop
	// return popped
	node := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return node
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	queries2 := make([]int, len(queries))
	copy(queries2, queries)
	sort.Ints(queries2)

	minHeap := &IntervalHeap{}
	heap.Init(minHeap)
	// book mark query with size solution
	ans := make(map[int]int)
	idx := 0

	for _, query := range queries2 {
		// push interval to min heap while query is still within range
		for ; idx < len(intervals) && query >= intervals[idx][0]; idx++ {
			heap.Push(minHeap, Interval{
				interval: intervals[idx],
				size:     intervals[idx][1] - intervals[idx][0] + 1,
			})
		}

		for minHeap.Len() > 0 && query > (*minHeap)[0].interval[1] {
			heap.Pop(minHeap)
		}

		if minHeap.Len() > 0 {
			ans[query] = (*minHeap)[0].size
		} else {
			ans[query] = -1
		}
	}

	ret := make([]int, len(queries))
	for idx, query := range queries {
		ret[idx] = -1
		if size, exists := ans[query]; exists {
			ret[idx] = size
		}
	}

	return ret
}
