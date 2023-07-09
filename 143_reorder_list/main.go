package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	log.Printf("debug 1 %v ", head)
	posNodeMap, length := getPosNodeMap(head)

	log.Printf("debug 2 %v %v", head, length)
	i, j := 0, length

	curr := head

	log.Printf("debug 3 %v ", posNodeMap[length])
	head.Next = posNodeMap[length]
	head = head.Next

	i++
	j--

	//log.Printf("debug 1 %v %v", i, j)
	//log.Printf("debug 2 %v ", posNodeMap)

	for i <= j {

		//log.Printf("debug 3 %v ", head)

		head.Next = posNodeMap[i]

		head = head.Next
		head.Next = posNodeMap[j]
		head = head.Next
		i++
		j--
	}

	head = curr
}

func getPosNodeMap(head *ListNode) (map[int]*ListNode, int) {
	curr := head
	posNodeMap := make(map[int]*ListNode)
	pos := 0
	for curr != nil {
		posNodeMap[pos] = curr
		curr = curr.Next
		pos++
	}
	return posNodeMap, pos - 1
}
