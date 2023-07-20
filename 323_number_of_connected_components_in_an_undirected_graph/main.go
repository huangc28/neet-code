package main

func countComponents(n int, edges [][]int) int {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	for _, edge := range edges {
		p1 := Find(edge[0], parents)
		p2 := Find(edge[1], parents)

		if p1 == p2 {
			continue
		}

		// they are in different components. assign p2's parent to be p1
		parents[p2] = p1
	}

	// Finally we need to cound number of unique root vertex. they are disjoint set
	roots := make(map[int]bool)
	for _, parent := range parents {
		p := Find(parent, parents)
		roots[p] = true
	}

	return len(roots)
}

func Find(vertex int, parents []int) int {
	parent := parents[vertex]
	for parent != parents[parent] {
		parent = parents[parent]
	}
	parents[vertex] = parent
	return parent
}
