package main

import "log"

/*
[100,4,200,1,3,2]

這些數字中，哪些數字是沒有 left neighbor? 這些數字就是 start of the sequence.
然後 iterate 這串 array numbr by numbr 檢查哪一些數字是 start of the sequence。
如果是 start of the sequence 我們就建立一個新的 sequence.
*/
func longestConsecutive(nums []int) int {
	numberMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		numberMap[nums[i]] = true
	}

	log.Printf("debug 1 %v", numberMap)

	seqs := make([][]int, 0)

	for j := 0; j < len(nums); j++ {
		num := nums[j]
		leftNeighbor := num - 1
		log.Printf("debug 2 %v", leftNeighbor)
		if _, hasLeftNeibor := numberMap[leftNeighbor]; hasLeftNeibor {
			continue
		} else {
			log.Printf("debug 3 %v", num)
			seq := make([]int, 0)
			seq = append(seq, num)

			nextSeq := num + 1
			for numberMap[nextSeq] {
				seq = append(seq, nextSeq)
				nextSeq += 1
			}
			seqs = append(seqs, seq)
		}
	}

	log.Printf("seqs %v", seqs)

	return 0
}
