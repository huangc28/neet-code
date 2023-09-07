package main

import "sort"

/*
lastSeen=1,2
count=4

[1,1,1,2,2,2,2]

	i

- 3/7 < 0.5
- 4/7 > 0.5

Time: O(N)
Space: O(1)
*/
func majorityElement(nums []int) int {
	sort.Ints(nums)

	lastSeen := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		// if current number is last seen we increase count by 1
		if nums[i] == lastSeen {
			count++
		} else {
			// If current number is not last seen, we calculate the frequency rate of last seen number
			// if its greater than 0.5, simply return the number
			if float64(count)/float64(len(nums)) > 0.5 {
				return lastSeen
			}

			// we'll set the lastSeen to be the current number.
			// set number count to be 0 since we're starting from a new character.
			lastSeen = nums[i]
			count = 1
		}
	}

	return lastSeen
}
