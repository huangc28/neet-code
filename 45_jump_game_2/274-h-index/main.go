package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	hIdx := len(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= hIdx {
			return hIdx
		}
		hIdx--
	}
	return 0
}

func hIndex_2(citations []int) int {
	totalPapers := len(citations)

	// Count number of publication that have 'n' citations
	// for example:
	//  - we have 1 publication with 1 citation.
	//  - we have 2 publications with 2 citations.
	//  - ...etc
	citationCounts := make([]int, len(citations)+1)

	for _, c := range citations {
		if c >= totalPapers {
			citationCounts[totalPapers] += 1
		} else {
			citationCounts[c] += 1
		}
	}

	// Ask ourself a question, how many publications
	// has at least `n` citations. If there are 3
	// publications has at least 3 citations, 3 is our h index.
	// [1,0,0,1,0,2]
	//  0,1,2,3,4,5 ---> number of publications
	h := 0
	for i := totalPapers; i >= 0; i-- {
		h += citationCounts[i]
		if h >= i {
			return i
		}
	}

	return 0
}
