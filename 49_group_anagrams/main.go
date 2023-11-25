package main

func getKey(str string) [26]int {
	keyCnt := [26]int{}
	for i := 0; i < len(str); i++ {
		keyCnt[str[i]-'a'] += 1
	}
	return keyCnt
}

func groupAnagrams(strs []string) [][]string {
	charGroup := make(map[[26]int][]string)

	for _, str := range strs {
		key := getKey(str)
		charGroup[key] = append(charGroup[key], str)
	}

	res := make([][]string, 0)
	for _, group := range charGroup {
		res = append(res, group)
	}

	return res
}
