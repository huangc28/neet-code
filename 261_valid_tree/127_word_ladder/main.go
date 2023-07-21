package main

import "log"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if wordList[len(wordList)-1] != endWord {
		return 0
	}

	// add beginWord to the front of the list
	wordList = append([]string{beginWord}, wordList...)

	// initialize parents array for union find for disjoin set
	// each word is  unique set
	parents := make([]int, len(wordList))
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	// start performing union find.
	left := 0
	right := left + 1
	for right < len(wordList) {
		p1 := findRootParent(left, parents)
		p2 := findRootParent(right, parents)

		// perform union when
		//  - 2 words have at most 1 char diff
		//  - have different root parent
		log.Printf("debug 1 %v, %v %v", wordList[left], wordList[right], hasSingleLetterDiff(wordList[left], wordList[right]))
		if hasSingleLetterDiff(wordList[left], wordList[right]) && p1 != p2 {
			parents[right] = left
			left = right
			right++

		} else {
			right++
		}
	}

	//log.Printf("debug 1 %v", parents)

	trfCount := 0
	for i := 1; i < len(parents); i++ {
		if findRootParent(i, parents) == 0 {
			trfCount++
		}
	}

	return trfCount
}

func findRootParent(vertex int, parents []int) int {
	parent := parents[vertex]
	for parent != parents[parent] {
		parent = parents[parent]
	}
	return parent
}

func hasSingleLetterDiff(word1, word2 string) bool {
	longerWord, shorterWord := word1, word2
	if len(longerWord) < len(shorterWord) {
		longerWord, shorterWord = shorterWord, longerWord
	}

	diffCount := 0
	i := 0
	for ; i < len(shorterWord); i++ {
		if longerWord[i] != shorterWord[i] {
			diffCount++
		}

		if diffCount > 1 {
			return false
		}
	}

	if len(longerWord)-1 > i {
		return false
	}

	return true
}
