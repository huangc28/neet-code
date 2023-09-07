package main

import (
	"container/heap"
	"sort"
)

/*
[1,2,3,6,2,3,4,7,8] groupSize = 3


[1,2,2,3,3,4,6,7,8]

- We have 9 card here, 9%3 = 0 we can possibly group them into group of 3
- Sort the card in ascending order
- So we have 3 rounds of cards to pick.
- How to pick cards? can we pick all cards within 3 rounds?
    - since card value >= 0, we can set the curr_card value to -1
    - when we encounter a card that we have not visited and it's value > curr_card value
      we mark that card as visited meaning that this card can be picked during this round
- In the end, we check if the length of the visited map equals length of hand, we consider
  the current hand can be rearranged into group of "groupSize"

*/

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)

	// How many rounds we need to pick
	rounds := len(hand) / groupSize

	// Record the index of card in hand that we've picked before.
	// we don't repicked the card in the next round.
	pickedMap := make(map[int]bool)

	for round := 0; round < rounds; round++ {
		numPickedAtThisRound, lastCardVal := 0, -1
		// Iterate through hand to start picking card
		for idx, card := range hand {
			// Only pick card that we haven't visited and it's value is greater than
			// the lastCard value
			if _, hasPicked := pickedMap[idx]; hasPicked {
				continue
			}

			// This card can be picked
			if lastCardVal == -1 || card-1 == lastCardVal {
				pickedMap[idx] = true
				lastCardVal = card
				numPickedAtThisRound++
			}

			// If we have picked `groupSize` card, quit this round
			// to start picking next round
			if numPickedAtThisRound == groupSize {
				break
			}
		}
	}

	return len(pickedMap) == len(hand)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return n
}
func (h *MinHeap) Push(v interface{}) {
	(*h) = append((*h), v.(int))
}

func isNStraightHand_2(hand []int, groupSize int) bool {
	freqMap := make(map[int]int)
	for _, card := range hand {
		freqMap[card]++
	}

	// Heapify hand
	minHeap := MinHeap(hand)
	heap.Init(&minHeap)

	for minHeap.Len() > 0 {
		minCard := minHeap[0]

		if freqMap[minCard] > 0 {
			for i := minCard; i < minCard+groupSize; i++ {
				if count, exists := freqMap[i]; count <= 0 || !exists {
					return false
				}
				freqMap[i]--
			}
		}

		// If we have used up all available card, pop the card from the heap.
		if freqMap[minCard] == 0 {
			heap.Pop(&minHeap)
		}
	}

	return true
}
