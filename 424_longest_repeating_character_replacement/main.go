package main

/*
Example: ABAB, k = 2


Example: A A B A B B A, k = 2

在目前 string 的長度裡面，我們只在乎出現最頻繁的 character。
我們只要 len(string) - (number of most freq char) = (number of times we need to replace)

要是 (number of times we need to replace) > k。代表我們要縮短 string 長度
- left = + 1
- 因為 left 往前移動，出現的次數也會減少 1，所以 charFreqMap[s[left]] -= 1
*/

func characterReplacement(s string, k int) int {
	charFreqMap := make(map[byte]int)
	left := 0
	maxFreq := 0
	for right := 0; right < len(s); right++ {
		charFreqMap[s[right]] += 1
		numOfMostFreqChar := findMaxInMap(charFreqMap)
		if ((right+1)-left)-numOfMostFreqChar > k {
			left += 1
			charFreqMap[s[left]] -= 1
		}

		maxFreq = max(maxFreq, (right+1)-left)
	}

	return maxFreq
}

func findMaxInMap(freqMap map[byte]int) int {
	maxCount := 0
	for _, v := range freqMap {
		maxCount = max(maxCount, v)
	}

	return maxCount
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
