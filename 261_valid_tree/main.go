package main

/*
- a valid tree should have only 1 root node
- a valid tree should not have cycle
*/
func validTree(n int, edges [][]int) bool {
	parents := make([]int, n)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	for _, edge := range edges {
		// check if they are in the same set
		// by checking they have the same root parent.
		// if they are, adding this edge would form a cycle
		p1 := find(edge[0], parents)
		p2 := find(edge[1], parents)

		if p1 == p2 {
			return false
		}

		// perform union
		parents[p2] = p1
	}

	roots := make(map[int]bool)
	for _, parent := range parents {
		p := find(parent, parents)
		roots[p] = true
	}

	if len(roots) > 1 {
		return false
	}
	return true
}

func find(vertex int, parents []int) int {
	parent := parents[vertex]
	for parent != parents[parent] {
		parent = parents[parent]
	}
	return parent
}
