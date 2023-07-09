package main

/*
s1: ab

s2: eidbaooo

Transform s1 to map:

{ a: true, b: true }

Iterate through s2. If we found a character in s2 that is in our s1 map, we use that position as left and right anchor.
Increase right by 1, if we found a match in s1 char map, mark it as true in visited map to indicate it's in the permutation until
we reach a chracter that is not in the visited map or it is already in the visited map.
*/
func checkInclusion(s1 string, s2 string) bool {
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
		} else {
			matches += 0
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
