package main

import (
	"log"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	// construct adj list
	adj := make(map[string][]string)
	for _, ticket := range tickets {
		adj[ticket[0]] = append(adj[ticket[0]], ticket[1])
	}

	// sort the each adj list in lexical order. so that we can traverse
	// the airport in the correct order.
	for _, ticket := range tickets {
		sort.Strings(adj[ticket[0]])
	}

	log.Printf("debug %v", adj)

	ans := make([]string, 0)

	var dfs func(from string) bool
	dfs = func(from string) bool {
		// If we've traveled all airports return true
		// vertex = ticket+1
		if len(ans) == len(tickets)+1 {
			return true
		}

		// if we have not finish traverse the airports
		// and the current airport has no outgoing airports
		// we return false
		if len(adj[from]) == 0 {
			return false
		}

		// we'll try travel to the adjacent airports
		for idx, airport := range adj[from] {
			// Try travel to the current adj airport
			ans = append(ans, airport)

			// We remove the current airport from the adj list because we don't want to
			// travel to this airport again.
			adj[from] = append(adj[from][:idx], adj[from][idx+1:]...)
			if dfs(airport) {
				return true
			}

			// Backtracking. We can not finish traveling all airport going this route.
			// add back the traveled airport and pop the current airport from answer.
			adj[from] = append(adj[from][:idx+1], adj[from][idx:]...)
			adj[from][idx] = airport
			ans = ans[:len(ans)-1]
		}

		return false
	}

	ans = append(ans, "JFK")
	dfs("JFK")

	return ans
}
