package main

func subsets(nums []int) [][]int {
	supersets := make([][]int, 0)
	supersets = append(supersets, []int{})

	for i := 0; i < len(nums); i++ {
		copysets := make([][]int, len(supersets))
		copy(copysets, supersets)
		for j := 0; j < len(copysets); j++ {
			copysets[i] = append([]int{nums[i]}, copysets[i]...)
		}
		supersets = append(supersets, copysets...)
	}

	return supersets
}
