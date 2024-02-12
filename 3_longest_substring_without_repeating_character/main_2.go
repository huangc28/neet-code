package main

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := make(map[byte]int)

	l, maxLen := 0, 1
	r := 1

	for r < len(s) {
		if seen[s[r]] > 0 {
			// deduct seen[s[l]] by 1
			seen[s[l]] -= 1
			// move l forward
			l += 1
		} else {
			// mark seen[s[r]] to be seen by increase count
			seen[s[r]] += 1

			maxLen = max(maxLen, r-l+1)

			// move r forward
			r += 1
		}
	}

	return maxLen
}
