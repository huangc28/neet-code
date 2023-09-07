package main

/*
1 word: abc, 3 patterns
2 word abd

*bc: [abc]
a*c: [abc]
ab*: [abc]

add each new word in the list,

number of edges in the tree (n)
*/
type WordInfo struct {
	Word string
	Dis  int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	adjMap := make(map[string][]string)
	for i, word := range wordList {
		pattern := word[:i] + "*" + word[i+1:]
		adjMap[pattern] = append(adjMap[pattern], word)
	}

	q := make([]WordInfo, 0)
	q = append(q, WordInfo{Word: beginWord, Dis: 1})
	visited := make(map[string]bool)

	for len(q) > 0 { // O(n-1)
		qLen := len(q)
		for i := 0; i < qLen; i++ { // O(n)
			wordInfo := q[0]
			q = q[1:]

			if wordInfo.Word == endWord {
				return wordInfo.Dis
			}

			visited[wordInfo.Word] = true

			for i := 0; i < len(wordInfo.Word); i++ { // O(m) number of characters in word
				pattern := wordInfo.Word[:i] + "*" + wordInfo.Word[i+1:]
				for _, adjWord := range adjMap[pattern] {
					if !visited[adjWord] {
						q = append(q, WordInfo{Word: adjWord, Dis: wordInfo.Dis + 1})
					}
				}
			}
		}
	}

	return 0
}
