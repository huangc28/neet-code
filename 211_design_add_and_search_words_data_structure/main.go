package main

type Trie struct {
	dict   map[byte]*Trie
	isWord bool
}

type WordDictionary struct {
	root *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Trie{dict: make(map[byte]*Trie)}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.dict[char]; !exists {
			curr.dict[char] = &Trie{dict: make(map[byte]*Trie)}
		}
		curr = curr.dict[char]
	}
	curr.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this.root
	var dfs func(i int, trie *Trie) bool
	dfs = func(i int, trie *Trie) bool {
		if i == len(word) {
			return trie.isWord
		}

		for j := i; i < len(word); i++ {
			char := word[i]
			if char == '.' {
				for v := range curr.dict {
					if dfs(j+1, curr.dict[v]) {
						return true
					}
				}
				return false
			} else {
				if _, exists := curr.dict[char]; !exists {
					return false
				}
				curr = curr.dict[char]
			}
		}

		return curr.isWord
	}

	return dfs(0, curr)
}
