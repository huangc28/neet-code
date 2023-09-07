package main

func countSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		oddCount := getOddPalCount(i, s)
		evenCount := getEvenPalCount(i, s)
		count += oddCount + evenCount
	}

	return count
}

func getOddPalCount(i int, s string) int {
	left, right, count := i, i, 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

func getEvenPalCount(i int, s string) int {
	left, right, count := i-1, i, 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
