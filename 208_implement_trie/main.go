package main

type Trie struct {
	dict   map[byte]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{dict: make(map[byte]*Trie)}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.dict[char]; !exists {
			curr.dict[char] = &Trie{dict: make(map[byte]*Trie)}
		}
		curr = curr.dict[char]
	}
	curr.isWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.dict[char]; !exists {
			return false
		}
		curr = curr.dict[char]
	}
	return curr.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, exists := curr.dict[char]; !exists {
			return false
		}
		curr = curr.dict[char]
	}
	return true
}
