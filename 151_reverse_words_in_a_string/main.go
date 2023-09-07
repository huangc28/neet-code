package main

import "strings"

func reverseWords(s string) string {
	words := make([]string, 0)
	word := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == ' ' {
			if string(word) != "" {
				words = append(words, string(word))
				word = make([]byte, 0)
			}
		} else {
			word = append(word, char)
		}
	}

	reverseWordsArr(words)
	return strings.Join(words, " ")
}

func reverseWordsArr(words []string) {
	l, r := 0, len(words)
	for l < r {
		words[l], words[r] = words[r], words[l]
	}
}
