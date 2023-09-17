package main

func partitionLabels(s string) []int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char := s[i]
		m[char] = i
	}

	start, end := 0, 0
	sizes := make([]int, 0)
	for i := 0; i < len(s); i++ {
		char := s[i]
		end = max(end, m[char])

		if i == end {
			sizes = append(sizes, end-start+1)
			end += 1
			start = end
		}
	}

	return sizes
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
