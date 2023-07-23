package main

type Trie struct {
	Children map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		Children: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.Children[char]; !exists {
			curr.Children[char] = &Trie{Children: make(map[byte]*Trie)}
		}
		curr = curr.Children[char]
	}
	curr.Children['*'] = nil
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.Children[char]; !exists {
			return false
		}
		curr = curr.Children[char]
	}
	_, end := curr.Children['*']
	return end
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, exists := curr.Children[char]; !exists {
			return false
		}
	}
	return true
}
