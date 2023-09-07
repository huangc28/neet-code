package main

import "log"

func numberOfSpecialSubstrings(s string) int {
	// We'll use a sliding window to get the count of substrings containing unique
	// characters
	var res int
	var counts [26]int
	left := 0
	for right := 0; right < len(s); right++ {
		ch := int(s[right] - 'a')
		counts[ch]++
		log.Printf("debug 1 %v %v", ch, counts[ch])
		// Make sure the window doesn't have any duplicates
		for counts[ch] > 1 {
			log.Printf("debug 2 %v", ch)

			ch2 := int(s[left] - 'a')
			counts[ch2]--
			left++
		}
		// We can add the length of window amount of new substrings that contain unique
		// characters.
		res += right - left + 1
	}
	return res
}
