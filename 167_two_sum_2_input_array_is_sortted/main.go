package main

/*
[2, 7, 11, 15]

	L          R
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) < 1 {
		return []int{-1, -1}
	}

	pLeft := 0
	pRight := len(numbers) - 1

	for pLeft < pRight {
		leftVal := numbers[pLeft]
		rightVal := numbers[pRight]
		sum := leftVal + rightVal

		if sum == target {
			return []int{pLeft + 1, pRight + 1}
		}

		if sum < target {
			pLeft += 1
		} else {
			pRight -= 1
		}
	}

	return []int{-1, -1}
}
