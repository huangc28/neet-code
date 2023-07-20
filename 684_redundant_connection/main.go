package main

func findRedundantConnection(edges [][]int) []int {
	// every node is a unique set at first.
	parents := make([]int, len(edges)+1)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}
	sizes := make([]int, len(edges)+1)
	for i := 0; i < len(sizes); i++ {
		sizes[i] = 1
	}

	for _, edge := range edges {
		p1 := find(edge[0], parents)
		p2 := find(edge[1], parents)

		if p1 == p2 {
			return edge
		}

		// perform union
		parents[p2] = p1
		sizes[p1] += sizes[p2]
	}
	return []int{}
}

func find(vertex int, parents []int) int {
	parent := parents[vertex]
	for parent != parents[parent] {
		parent = parents[parent]
	}
	// collasping find
	parents[vertex] = parent
	return parent
}
