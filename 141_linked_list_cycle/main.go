package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	f := head.Next
	s := head

	for s != f {
		if f == nil {
			return false
		}

		// s move 1 step
		s = s.Next

		// f move 2 steps
		f = f.Next
		if f == nil {
			return false
		}
		f = f.Next
	}

	return true
}
