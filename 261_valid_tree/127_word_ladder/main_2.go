package main

/*
create pattern adjacency list:

	{
		"*it": [hit]
		"h*t": [hit]
		"hi*": [hit]
		"*ot": [dot]
		"d*t": [dot]
		"do*": [dot]
	}
*/
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if !Contains(endWord, wordList) {
		return 0
	}

	adjList := make(map[string][]string)
	wordList = append(wordList, beginWord)

	// create adjacency list
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			adjList[pattern] = append(adjList[pattern], word)
		}
	}

	// pop off current word from queue.
	// retrieve all words from all possible pattern. For example, *it:[...], h*t:[...], hi*:[...]
	// if we have visited the pattern, skip it
	// if current word equals endWord, we've found the sequence. return the shortest number of transformation sequence.
	visitedPattern := make(map[string]bool)
	visitedPattern[beginWord] = true
	distance := 1
	queue := make([]string, 0)
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			currWord := queue[0]
			queue = queue[1:]

			if currWord == endWord {
				return distance
			}

			// push all possible valid transformations to queue
			for j := 0; j < len(currWord); j++ {
				pattern := currWord[:j] + "*" + currWord[j+1:]
				for _, neiWord := range adjList[pattern] {
					if !visitedPattern[neiWord] {
						visitedPattern[neiWord] = true
						queue = append(queue, neiWord)
					}
				}
			}
		}
		distance += 1
	}

	return 0
}

func Contains(word string, wordList []string) bool {
	for _, wordInList := range wordList {
		if word == wordInList {
			return true
		}
	}
	return false
}
