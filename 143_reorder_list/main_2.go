package main

func reorderList2(head *ListNode) {
	length := countLength(head)

	if length == 1 {
		return
	}

	half := length / 2

	// This would be our first list
	curr := head
	for i := 1; i < half; i++ {
		curr = curr.Next
	}

	// This would be our second list
	sec := curr.Next

	// cut off the link between first and sec list
	curr.Next = nil

	// reverse the sec list
	sec = reverse(sec)

	// combine head and sec to new list.
	var (
		newList *ListNode
		newHead *ListNode
	)

	newHead = newList

	for head != nil && sec != nil {
		if newList == nil {
			newList = head
			newHead = head
		} else {
			newList.Next = head
			newList = newList.Next
		}

		head = head.Next

		newList.Next = sec
		newList = newList.Next
		sec = sec.Next
	}

	if head != nil {
		newList.Next = head
		newList = newList.Next
	}

	if sec != nil {
		newList.Next = sec
		newList = newList.Next
	}

	head = newHead
}

func reverse(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func countLength(head *ListNode) int {
	curr := head
	count := 0

	for curr != nil {
		curr = curr.Next
		count++
	}

	return count
}
