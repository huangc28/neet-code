package main

type WordDictionary struct {
	dict map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{dict: make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	var cb func(i int, dict *WordDictionary)
	cb = func(i int, wd *WordDictionary) {
		if i == len(word) {
			wd.dict['*'] = nil
			return
		}

		// add current byte char to dict at current position
		if _, exists := wd.dict[word[i]]; !exists {
			wd.dict[word[i]] = &WordDictionary{dict: make(map[byte]*WordDictionary)}
		}
		cb(i+1, wd.dict[word[i]])

		// add current char position to be .
		if _, exists := wd.dict['.']; !exists {
			wd.dict['.'] = &WordDictionary{dict: make(map[byte]*WordDictionary)}
		}
		cb(i+1, wd.dict['.'])
	}
	cb(0, curr)
}

func (this *WordDictionary) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := curr.dict[char]; !exists {
			return false
		}
		curr = curr.dict[char]
	}
	_, end := this.dict['*']
	return end
}
