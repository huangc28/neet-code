package main

/*
s = "ADOBECODEBANC"
t = "ABC"

matches: 3

{ A:1, B:1, C: 1 }

tCount = {A:1, B:1, C:1}
*/
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tCount := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tCount[t[i]] += 1
	}

	minSub := s
	left := 0
	matches := 0
	sCount := map[byte]int{}

	for right := 0; right < len(s); right++ {
		rightChar := s[right]

		// if current character is in `t`, add character count
		if _, exists := tCount[rightChar]; !exists {
			continue
		}

		sCount[rightChar] += 1
		if sCount[rightChar] == tCount[rightChar] {
			matches++
		}

		if matches < len(t) {
			continue
		}

		// we found a potensital matching min window
		// update current min window if current window
		// is smaller

		if minSub == "" {
			minSub = s[left : right+1]
		} else {
			minSub = minStr(minSub, s[left:right+1])
		}

		// move left pointer to right until matches != len(s)
		for matches == len(tCount) {
			minSub = minStr(minSub, s[left:right+1])
			leftChar := s[left]

			if _, exists := tCount[leftChar]; exists {
				sCount[leftChar]--
				if sCount[leftChar] < tCount[leftChar] {
					matches--
				}
			}

			left++
		}
	}

	return minSub
}

func minStr(str1, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	}
	return str2
}
