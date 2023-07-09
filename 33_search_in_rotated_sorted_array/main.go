package main

func search(nums []int, target int) int {
	pLeft := 0
	pRight := len(nums) - 1

	for pLeft <= pRight {
		pMid := (pLeft + pRight) / 2
		if nums[pMid] == target {
			return pMid
		}

		if nums[pMid] >= nums[pLeft] {
			if target >= nums[pLeft] && target < nums[pMid] {
				// explore left
				pRight = pMid - 1
			} else {
				// explore right
				pLeft = pMid + 1
			}
		} else {
			if nums[pRight] >= target && target > nums[pMid] {
				// explore right
				pLeft = pMid + 1
			} else {
				// explore left
				pRight = pMid - 1
			}
		}
	}

	return -1
}
