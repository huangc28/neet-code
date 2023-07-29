package main

/*
max heap

We start process tasks we most frequence in order to have least number of units of times

After we process each task, we push it to queue, when the next time
we finish the other task check if it is time to push task from top of the queue to
max heap.

*/

type TaskInfo struct {
	Freq     int
	NextTime int
}

func leastInterval(tasks []byte, n int) int {
	freqMap := make(map[byte]int)
	for _, task := range tasks {
		freqMap[task]++
	}
	taskFreqs := make([]int, 0, len(tasks))
	for _, freq := range freqMap {
		taskFreqs = append(taskFreqs, freq)
	}

	MaxHeapify(taskFreqs)
	queue := make([]TaskInfo, 0)
	time := 0
	for len(taskFreqs) > 0 || len(queue) > 0 {
		time++

		if len(taskFreqs) == 0 {
			time = queue[0].NextTime
		} else {
			taskFreq := pop(&taskFreqs) - 1
			if taskFreq > 0 {
				// still tasks left to be done push to queue
				queue = append(queue, TaskInfo{Freq: taskFreq, NextTime: time + n})
			}
		}

		if len(queue) > 0 && queue[0].NextTime == time {
			n := queue[0]
			queue = queue[1:]
			add(&taskFreqs, n.Freq)
		}
	}

	return time
}

func MaxHeapify(nums []int) {
	curr := len(nums) - 1
	for parent := (curr - 1) / 2; parent >= 0; parent-- {
		siftDown(nums, parent, len(nums)-1)
	}
}

func siftDown(nums []int, startIdx, endIdx int) {
	for startIdx <= endIdx {
		lc := (startIdx * 2) + 1
		rc := (startIdx * 2) + 2
		idxToSwap := -1

		if lc <= endIdx {
			idxToSwap = lc
			if rc <= endIdx && nums[rc] > nums[lc] {
				idxToSwap = rc
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

func siftUp(nums []int) {
	curr := len(nums) - 1
	for curr >= 0 {
		parent := (curr - 1) / 2
		if nums[curr] > nums[parent] {
			swap(nums, curr, parent)
			curr = parent
		} else {
			return
		}
	}
}

func add(nums *[]int, num int) {
	*nums = append(*nums, num)
	siftUp(*nums)
}

func pop(nums *[]int) int {
	swap(*nums, 0, len(*nums)-1)
	n := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	siftDown(*nums, 0, len(*nums)-1)

	return n
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
