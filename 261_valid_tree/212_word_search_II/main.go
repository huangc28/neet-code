package main

import "log"

type Trie struct {
	children map[byte]*Trie
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{children: make(map[byte]*Trie)}
	output := make([]string, 0)
	visited := make([][]bool, len(board))
	for y := range visited {
		visited[y] = make([]bool, len(board[y]))
	}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// build trie using dfs
	var dfs func(y, x int, root *Trie, visited [][]bool)
	dfs = func(y, x int, root *Trie, visited [][]bool) {
		if y < 0 ||
			x > len(board[0])-1 ||
			y > len(board)-1 ||
			x < 0 ||
			visited[y][x] {
			return
		}

		visited[y][x] = true
		curr := root
		char := board[y][x]
		if _, exists := curr.children[char]; !exists {
			curr.children[char] = &Trie{children: make(map[byte]*Trie)}
		}
		curr = curr.children[char]

		// traverse up, right, down, left recursively
		for _, dir := range directions {
			dfs(y+dir[0], x+dir[1], curr, visited)
		}
		visited[y][x] = false
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			dfs(0, 0, root, visited)
		}
	}

	log.Printf("debug 1 %v", root)

	for _, word := range words {
		if findWord(word, root) {
			output = append(output, word)
		}
	}

	return output
}

func findWord(word string, root *Trie) bool {
	curr := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, exists := curr.children[c]; !exists {
			return false
		}
		curr = curr.children[c]
	}
	return true
}
