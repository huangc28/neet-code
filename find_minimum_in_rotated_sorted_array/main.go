func fundMin(nums []int) int {
	pLeft := 0
	pRight := len(nums) - 1
	smallest := nums[0]

	for pLeft <= pRight {
		pMid := (pLeft + pRight) / 2
		if nums[pMid] < smallest {
			smallest = nums[pMid]
		}

		if nums[pMid] >= nums[0] {
			// search right
			pLeft = pMid + 1
		} else {
			// search left
			pRight = pMid - 1
		}
	}

	return smallest
}
