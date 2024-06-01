package main

// dummy node connect to head
// set prevGroup to dummy
// find kth node
// start reverse linked list until kth.next
// prevGroup.Next = kth
// firstNodeOfGroup.Next = nextGroup
// prevGroup = firstNodeOfGroup

func reverseKGroup_2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	prevGroup := dummy

	for {
		kthNode := findKthNode(prevGroup, k)

		if kthNode == nil {
			break
		}

		nextGroup := kthNode.Next
		prev, curr := prevGroup, prevGroup.Next

		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// 1. next node of previous group should be last node of current group
		// 2. first node's node of current group should be next group
		firstNodeOfGroup := prevGroup.Next
		prevGroup.Next = kthNode
		firstNodeOfGroup.Next = nextGroup
		prevGroup = firstNodeOfGroup
	}

	return dummy.Next
}

/*
1 -> 2 -> 3, k = ~2,~1,0
          c
*/

func findKthNode(head *ListNode, k int) *ListNode {
	curr := head
	for curr != nil && k > 0 {
		curr = curr.Next
		k -= 1
	}
	return curr
}
