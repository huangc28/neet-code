package main

/*
abcabcbb
*/
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]bool)
	maxCount, length := 0, 0
	left := 0

	for right := 0; right < len(s); right++ {
		if charMap[s[right]] {
			for ; charMap[s[right]]; left++ {
				length--
				charMap[s[left]] = false
			}
		}

		charMap[s[right]] = true
		length++
		maxCount = max(maxCount, length)
	}

	return 0
}

func lengthOfLongestSubstring_2(s string) int {
	charMap := make(map[byte]int)
	left, maxCount := 0, 0

	for right := 0; right < len(s); right++ {
		lastSeenPos, hasSeen := charMap[s[right]]

		if !hasSeen {
			charMap[s[right]] = right
			maxCount = max(maxCount, right-left+1)
			continue
		}

		newLastSeenPost := lastSeenPos + 1
		left = max(left, newLastSeenPost)
		maxCount = max(maxCount, right+1-left)
		charMap[s[right]] = right
	}

	return 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
