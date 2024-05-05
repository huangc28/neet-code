package main

func hasCycle_2(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
	}

	return false
}
