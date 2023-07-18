package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	visited := make(map[int]bool)

	for i := 0; i < len(prerequisites); i++ {
		prerequisite := prerequisites[i]
		adjList[prerequisite[0]] = append(adjList[prerequisite[0]], prerequisite[1])
	}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		// If this course has no other prerequisite courses
		if visited[course] {
			return false
		}
		visited[course] = true

		for _, pre := range adjList[course] {
			if !dfs(pre) {
				return false
			}
		}

		adjList[course] = []int{}
		visited[course] = false

		return true
	}

	for course := 0; course < numCourses; course++ {
		if !dfs(course) {
			return false
		}
	}
	return true
}
