package main

func maximumDetonation_2(bombs [][]int) int {
	bombsCanDetonate := make(map[int][]int)

	// create adjacency list
	// {
	//   0: [1, 2]
	//   1: [2, 4]
	//   3: [4, 5]
	// }
	for i := 0; i < len(bombs); i++ {
		for j := 0; j < len(bombs); j++ {
			if i != j {
				if canDetonate(bombs[i], bombs[j]) {
					bombsCanDetonate[i] = append(bombsCanDetonate[i], j)
				}
			}
		}
	}

	maxDetonateCount := 0

	for i := 0; i < len(bombs); i++ {
		queue := make([]int, 0)
		queue = append(queue, i)
		visited := make(map[int]bool)
		visited[i] = true

		for len(queue) > 0 {
			bombToStart := queue[0]
			queue = queue[1:]
			bombsToTest, exists := bombsCanDetonate[bombToStart]

			if exists {
				for _, bombToTest := range bombsToTest {
					if !visited[bombToTest] {
						visited[bombToTest] = true
						queue = append(queue, bombToTest)
					}
				}
			}
		}

		if len(visited) == len(bombs) {
			return len(visited)
		}

		maxDetonateCount = max(maxDetonateCount, len(visited))
	}

	return maxDetonateCount
}
