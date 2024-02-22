package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Freq := make(map[byte]int)
	s2Freq := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Freq[s1[i]-'a'] += 1
		s2Freq[s2[i]-'a'] += 1
	}

	// calculate number of matches of all 26 alphabets between s1 and s2
	matches := 0
	for j := 0; j < 26; j++ {
		if s1Freq[byte(j)] == s2Freq[byte(j)] {
			matches++
		}
	}

	// start sliding window on s2. window side should be the length of s1
	left := 0
	for right := len(s1); right < len(s2); right++ {
		// if the current window has all 26 char matches
		// return true.
		if matches == 26 {
			return true
		}

		// right has expanded by 1, increment count of s2[right]
		s2Freq[s2[right]-'a'] += 1

		// what if s2[right] was greater than s1[right]?  we don't need to increment matches.
		// only increment matches when s2[right]
		//
		// s1=aab
		// s2=aaccdaab
		//    l  r
		//
		if s2Freq[s2[right]-'a']-1 == s1Freq[s2[right]-'a']-1 {
			matches += 1
		} else {
			matches -= 1
		}

		// shrink left by 1, decrement count of s2[left]
		s2Freq[s2[left]-'a'] -= 1
		if s2Freq[s2[left]-'a']+1 == s1Freq[s2[left]-'a']+1 {
			matches -= 1
		}
		left += 1
	}

	return matches == 26
}
