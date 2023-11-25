package main

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int) // key: num, val: freq

	for _, num := range nums {
		freqMap[num] += 1
	}

	// create frequency bucket
	freqBuckets := make([][]int, len(nums)+1)

	// place num:freq to freqBucket
	for num, freq := range freqMap {
		freqBuckets[freq] = append(freqBuckets[freq], num)
	}

	// start counting top k freq numbers
	sol := make([]int, 0)

OUTER:
	for i := len(freqBuckets) - 1; i > 0; i-- {
		bucket := freqBuckets[i]

		if len(bucket) == 0 {
			continue
		}

		for _, num := range bucket {
			sol = append(sol, num)
			k -= 1
			if k == 0 {
				break OUTER
			}
		}
	}

	return sol
}
