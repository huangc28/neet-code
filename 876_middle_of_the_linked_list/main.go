package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode_2pointers(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next

		if fast.Next == nil {
			break
		}

		fast = fast.Next
	}

	return slow
}

func middleNode_count(head *ListNode) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	mid := length / 2

	for i := 0; i < len(mid); i++ {
	}
}
