package main

import (
	"log"
	"math"
)

/*
k closest points, if k = 3
[1, 2, 3, 4]

	k ---> k closest point

- use a min heap. calc and push distance to a min heap
- pop k times to find the closest point
*/
type Point struct {
	y, x int
	dis  float64
}

func kClosest(points [][]int, k int) [][]int {
	// initialize points
	ps := make([]*Point, len(points))
	for idx, point := range points {
		ps[idx] = &Point{
			x:   point[0],
			y:   point[1],
			dis: calcDistance(point[0], point[1]),
		}
	}

	// heapify points
	BuildMinHeap(&ps)

	res := make([][]int, 0)
	for i := 0; i < k; i++ {
		pt := pop(&ps)
		//log.Printf("debug 2 %v", pt)
		res = append(res, []int{pt.x, pt.y})

	}
	return res
}

func calcDistance(x, y int) float64 {
	xd := math.Pow(float64(x-0), 2)
	yd := math.Pow(float64(y-0), 2)
	return math.Sqrt(xd + yd)
}

func BuildMinHeap(pts *[]*Point) {
	curr := len(*pts) - 1
	for parent := (curr - 1) / 2; parent >= 0; parent-- {
		log.Printf("debug 5 %v %v", parent, len(*pts)-1)
		siftDown(*pts, parent, len(*pts)-1)
	}
}

func siftDown(pts []*Point, startIdx, endIdx int) {

	for startIdx <= endIdx {
		lc := startIdx*2 + 1
		rc := startIdx*2 + 2
		idxToSwap := -1

		if lc <= endIdx {
			idxToSwap = lc
			if rc <= endIdx && pts[rc].dis < pts[lc].dis {
				idxToSwap = rc
			}
		}

		if idxToSwap != -1 && pts[idxToSwap].dis < pts[startIdx].dis {
			//log.Printf("debug 6 %v", idxToSwap)
			swap(pts, idxToSwap, startIdx)
			startIdx = idxToSwap
		} else {
			return
		}
	}
}

func pop(pts *[]*Point) *Point {
	swap(*pts, 0, len(*pts)-1)
	pt := (*pts)[len(*pts)-1]
	(*pts) = (*pts)[:len(*pts)-1]
	siftDown(*pts, 0, len(*pts)-1)
	return pt
}

func swap(pts []*Point, i, j int) {
	pts[i], pts[j] = pts[j], pts[i]
}
