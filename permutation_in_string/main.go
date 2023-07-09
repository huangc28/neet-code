package main

/*
s1: ab

s2: eidbaooo

當 matches 為 26 代表 s1 的 permutation 存在於 s2。當我們在檢查 s2 時，我們一次只要考慮 s1 的長度。

每次 iterate slide left 跟 right 一直保持 s1 的長度。記得要調整 matches 的數字跟 s1 & s2 的 character count。當 s1 跟 s2 的那個  character count 不同，matches -= 1。要是相同 matches += 1。


abc
bbbca

matches: 24

{
 a:1
 b:1
 c:1
}

{
 a:1
 b:3,2
 c:1
}

1.
 abc
 bbb

2.
 abc
 bbc


*/

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	arr1, arr2 := [26]int{}, [26]int{}

	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a'] += 1
		arr2[s2[i]-'a'] += 1
	}

	matches := 0

	for i := 0; i < 26; i++ {
		if arr1[i] == arr2[i] {
			matches++
		}
	}

	left := 0
	for right := len(s1); right < len(s2); right++ {
		if matches == 26 {
			return true
		}

		index := s2[right] - 'a'

		// expand right by 1
		arr2[index] += 1
		if arr2[index] == arr1[index] {
			matches += 1
		} else if arr1[index] == arr2[index]-1 {
			matches -= 1
		}

		// shrink left by 1
		index = s2[left] - 'a'
		arr2[index] -= 1
		if arr2[index] == arr1[index] {
			matches += 1
		} else if arr1[index] == arr2[index]+1 {
			matches -= 1
		}
		left += 1
	}

	return matches == 26
}
