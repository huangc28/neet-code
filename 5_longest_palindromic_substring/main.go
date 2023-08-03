package main

/*
For every character, we treat it as center and expand outward.
If both characters are the same they are palindrome, if not they are not palindrome

An edge case would be checking odd palindrome: treat both s[i], s[i-1] are center and expand
outward simultaneously.
*/
func longestPalindrome(s string) string {
	maxPal := ""
	for i := 0; i < len(s); i++ {
		// treat s[i] as center and expand outward
		oddPal := getOddPal(i, s)

		// treat s[i] & s[i-1] as center and expand outward
		evenPal := getEvenPal(i, s)

		// which one is longer?
		longer := maxStr(oddPal, evenPal)

		// which one is longer the longer pal or max pal?
		maxPal = maxStr(maxPal, longer)

	}
	return maxPal
}

func getOddPal(idx int, s string) string {
	left, right := idx-1, idx+1
	for left >= 0 && right < len(s) && s[left] == s[right] {
		right++
		left--
	}
	return s[left+1 : right]
}

func getEvenPal(idx int, s string) string {
	left, right := idx-1, idx
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func maxStr(i, j string) string {
	if len(i) > len(j) {
		return i
	}
	return j
}
