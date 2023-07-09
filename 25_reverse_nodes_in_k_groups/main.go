package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prevGroup := dummy

	for {
		kthNode := getKth(prevGroup, k)
		if kthNode == nil {
			break
		}

		nextGroup := kthNode.Next
		prev, curr := prevGroup, prevGroup.Next

		// reverse k group list
		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// relink prev group next node to be the last node of reversed list
		firstNodeOfGroup := prevGroup.Next
		prevGroup.Next = kthNode
		firstNodeOfGroup.Next = nextGroup
		prevGroup = firstNodeOfGroup
	}

	return dummy.Next
}

func getKth(head *ListNode, k int) *ListNode {
	curr := head
	for curr != nil && k > 0 {
		curr = curr.Next
		k -= 1
	}
	return curr
}
