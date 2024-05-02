package main

func removeNthFromEnd_2(head *ListNode, n int) *ListNode {
	curr := head
	var rp *ListNode
	for i := 0; i < n; i++ {
		curr = curr.Next
	}
	rp = curr

	dummy := &ListNode{Val: -1, Next: head}
	lp := dummy

	for rp != nil {
		rp = rp.Next
		lp = lp.Next
	}

	lp.Next = lp.Next.Next

	return dummy.Next
}
