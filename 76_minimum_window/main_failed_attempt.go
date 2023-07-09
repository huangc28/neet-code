package main

import "log"

/*
s = "ADOBECODEBANC"
t = "ABC"

matches := 3
*/

func minWindow_failed(s string, t string) string {
	tmap := map[byte]bool{}
	leftPosList := make([]int, 0, len(t))

	// compose character map of t so we can compare it with character in s
	for i := 0; i < len(t); i++ {
		tchar := t[i]
		tmap[tchar] = true
	}

	// compose list of left positions
	for i := 0; i < len(s); i++ {
		schar := s[i]
		if _, exists := tmap[schar]; exists {
			leftPosList = append(leftPosList, i)
		}
	}

	log.Printf("debug 1 %v", leftPosList)

	matches := 0
	minSub := s
	sCharCount := make(map[byte]int)

	leftPointer := 0
	left := leftPosList[leftPointer]

	if tmap[s[left]] {
		sCharCount[s[left]] += 1
		matches += 1
	}

	for right := left + 1; right < len(s); right++ {
		currMinSubCount := right - left

		rightChar := s[right]

		log.Printf("debug 2 %v", string(rightChar))
		if _, exists := tmap[rightChar]; exists {
			if count, exists := sCharCount[rightChar]; !exists || count == 0 {
				matches += 1
			}
			sCharCount[rightChar] += 1
		}

		currMinSubCount++

		// we found a matching substring from left to right
		if matches == len(t) {

			log.Printf("debug 3 %v", s[left:right+1])
			if currMinSubCount < len(minSub) {
				minSub = s[left : right+1]
			}

			leftChar := s[left]

			// decrement the left character count by 1
			sCharCount[leftChar] -= 1

			// if that character count reaches 0, decrement matches
			if sCharCount[leftChar] == 0 {
				matches -= 1
			}

			// Move left to next position
			leftPointer += 1
			left = leftPosList[leftPointer]

			// and we check again...
			log.Printf("debug 4 %v", s[left:right+1])
			if len(s[left:right+1]) >= len(t) && matches == len(t) {
				if len(s[left:right+1]) < len(minSub) {
					minSub = s[left : right+1]
				}
			}
		}

	}

	return minSub
}
