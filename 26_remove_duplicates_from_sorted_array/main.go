package main

func removeDuplicates(nums []int) int {
	haveSeenMap := make(map[int]bool)
	pos := 0
	for _, num := range nums {
		if haveSeenMap[num] {
			continue
		}

		nums[pos] = num
		haveSeenMap[num] = true
		pos++
	}
	return pos
}

func removeDuplicates_2(nums []int) int {
	lastSeen := nums[0]
	replacePos := 1
	length := 1

	for i := 1; i < len(nums); i++ {
		// If we've seen current number in the previous step, continue checking
		// next duplicated number
		if nums[i] == lastSeen {
			continue
		}

		// If current number hasn't been seen (distinct number), replace this number
		// with current position
		nums[i] = lastSeen
		lastSeen = nums[i]
		replacePos++
		length++
	}

	return length
}
