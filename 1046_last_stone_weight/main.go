package main

/*
arr = [2, 7, 4, 1, 8 , 1]
[7, 8] result in [2, 4, 1, 1, 1]
[2, 4] result in [2, 1, 1, 1]
[2, 1] result in [1, 1, 1]
[1, 1] result in [1]

heapify arr. if arr length is > 1, we pop 2 numbers from heap.
those number will be the top 2 numbers in the arr (x, y).

if x == y, we don't add left over weight to max heap
if x < y, we add (y - x) to heap

when length heap <= 1 return max heap, that'll be our solution
*/

func lastStoneWeight(stones []int) int {
	buildMaxHeap(stones)
	for len(stones) > 1 {
		w1 := pop(&stones)
		w2 := pop(&stones)

		if w1 == w2 {
			continue
		}

		add(&stones, abs(w1-w2))
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

func buildMaxHeap(nums []int) {
	curr := len(nums) - 1
	for parent := (curr - 1) / 2; parent >= 0; parent-- {
		siftDown(nums, parent, len(nums)-1)
	}
}

func pop(nums *[]int) int {
	swap(*nums, 0, len(*nums)-1)
	num := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	siftDown((*nums), 0, len(*nums)-1)
	return num
}

func add(nums *[]int, num int) {
	*nums = append((*nums), num)
	siftUp(*nums)
}

func siftUp(nums []int) {
	curr := len(nums) - 1
	for curr > 0 {
		parent := (curr - 1) / 2
		if nums[curr] > nums[parent] {
			swap(nums, curr, parent)
			curr = parent
		} else {
			return
		}
	}
}

func siftDown(nums []int, startIdx, endIdx int) {
	for startIdx <= endIdx {
		leftChild := (startIdx * 2) + 1
		rightChild := (startIdx * 2) + 2
		idxToSwap := -1

		if leftChild <= endIdx {
			idxToSwap = leftChild
			if rightChild <= endIdx && nums[rightChild] > nums[leftChild] {
				idxToSwap = rightChild
			}
		}

		if idxToSwap != -1 && nums[idxToSwap] > nums[startIdx] {
			swap(nums, idxToSwap, startIdx)
			startIdx = idxToSwap
		} else {
			return
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func abs(i int) int {
	if i < 0 {
		return i * (-1)
	}
	return i
}
