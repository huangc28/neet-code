package main

type Trie struct {
	dict map[byte]*Trie
	word string
}

func (root *Trie) Insert(word string) {
	curr := root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.dict[char]; !exists {
			curr.dict[char] = &Trie{dict: make(map[byte]*Trie)}
		}
		curr = curr.dict[char]
	}
	curr.word = word
}

// direction vector (y,x) up, right, down, left
var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{dict: make(map[byte]*Trie)}
	for _, word := range words {
		root.Insert(word)
	}

	ans := make([]string, 0)
	visited := make([][]bool, len(board))
	for i := range board {
		visited[i] = make([]bool, len(board[i]))
	}

	var dfs func(y, x int, trie *Trie)
	dfs = func(y, x int, trie *Trie) {
		// if current coordinate has been visited or out of bound quick checking
		if y > len(board)-1 || y < 0 || x > len(board[0])-1 || x < 0 || visited[y][x] {
			return
		}

		// if current char at y,x does not exists in trie, return.
		char := board[y][x]
		charTrie, exists := trie.dict[char]
		if !exists {
			return
		}

		// if char trie exists and is a complete word, push this word to `ans`
		if charTrie.word != "" {
			ans = append(ans, charTrie.word)
			charTrie.word = ""
		}

		// mark current position as visited so we don't revisit the current
		// position in the next iteration
		visited[y][x] = true

		// if char exists in trie but is not a complete word yet,
		// we keep looking for next char in the adj positions recursivly.
		for _, dir := range dirs {
			dfs(y+dir[0], x+dir[1], charTrie)
		}

		// mark current position as unvisited
		visited[y][x] = false
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			curr := root
			dfs(y, x, curr)
		}
	}

	return ans
}
