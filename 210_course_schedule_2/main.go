package main

/*
[[1,0][2,0][3,1][3,2]]

	{
	 0: []
	 1: [0]
	 2: [0]
	 3: [1, 2]
	}

cycle: {}
visited: {0,1,2,3}

output: []
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	for _, pre := range prerequisites {
		adjList[pre[0]] = append(adjList[pre[0]], pre[1])
	}

	res := make([]int, 0)
	visited, cycle := make(map[int]bool), make(map[int]bool)
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if cycle[course] {
			return false
		}

		if visited[course] {
			return true
		}

		for _, precourse := range adjList[course] {
			if !dfs(precourse) {
				return false
			}
		}
		cycle[course] = false
		visited[course] = true
		res = append(res, course)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return res
}
