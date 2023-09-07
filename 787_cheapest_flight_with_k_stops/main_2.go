package main

import "math"

type Node struct {
	dstIdx int
	cost   int
	stops  int
}

/*

0 ---> 2
我一開始挑選 cost 比較少的 stop 1 走。這個 stop 有兩條路可以走。這時 distance[1] = 4。我們發覺這條路走不了。


*/

func findCheapestPrice_shortest_distance(n int, flights [][]int, src int, dst int, k int) int {
	adjMap := make(map[int][]Node)
	for _, flight := range flights {
		adjMap[flight[0]] = append(adjMap[flight[0]], Node{
			dstIdx: flight[1],
			cost:   flight[2],
		})
	}

	// storing minimum cost to reach the stop.
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt32
	}

	queue := make([]Node, 0)
	queue = append(queue, Node{
		dstIdx: src,
		cost:   0,
		stops:  0,
	})

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if currNode.stops > k {
			continue
		}

		for _, adjNode := range adjMap[currNode.dstIdx] {
			nextCost := adjNode.cost + currNode.cost
			if nextCost < distance[adjNode.dstIdx] {
				queue = append(queue, Node{
					dstIdx: adjNode.dstIdx,
					cost:   nextCost,
					stops:  currNode.stops + 1,
				})
				distance[adjNode.dstIdx] = nextCost
			}
		}
	}

	if distance[dst] == math.MaxInt32 {
		return -1
	}

	return distance[dst]
}
