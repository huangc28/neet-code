package main

import "log"

/*
words = ["wrt","wrf","er","ett","rftt"]

	l     r

t:[f]
w:[e]
r:[t]
e:[r]

wrt, wrf
't' & 'f' are different,
wrf, er
w & e
er, ett
r & t
ett, rftt
e & r

dfs traverse
[f,t,r,e,w] reverse [w,e,r,t,f]

neighbors = adjList['t']

cycle   = []
visited = [f,t,r,e,w]
ans     = [f,t,r,e,w]

	  f(t)
	  /
	f(f)

	  f(w)
	  /
	f(e)
	  |
	f(r)
	  |
	f(t)

traverse each neighbors recursively.
  - If char has been visited, skip
  - If char has cycle, return false
  - if no cycle and not visited yet, keep traverse it's neighbors
  - If done traversing neighbors, add itself to solution, mark it as visited, pop itself from cycle
*/
func alienOrder(words []string) string {
	adjList := make(map[byte][]byte)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			adjList[word[i]] = make([]byte, 0)
		}
	}

	left := 0
	for right := left + 1; right < len(words); right++ {
		lw, rw := words[left], words[right]

		if len(lw) > len(rw) && lw[:len(rw)] == rw {
			return ""
		}

		// find the first different char in from both words
		i := 0
		for i < len(lw) && i < len(rw) {
			if lw[i] != rw[i] {
				adjList[lw[i]] = append(adjList[lw[i]], rw[i])
				break
			}
			i++
		}

		left++
	}

	visited := make(map[byte]bool)
	cycle := make(map[byte]bool)
	ans := []byte{}

	var dfs func(char byte) bool
	dfs = func(char byte) bool {
		log.Printf("debug 3 %v", string(char))

		if visited[char] {
			return true
		}

		if cycle[char] {
			return false
		}

		cycle[char] = true

		for _, neighborChar := range adjList[char] {
			if !dfs(neighborChar) {
				return false
			}
		}

		cycle[char] = false
		visited[char] = true
		ans = append(ans, char)
		return true
	}

	for char := range adjList {
		if !dfs(char) {
			return ""
		}
	}

	return reverse(ans)
}

func reverse(ans []byte) string {
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}
