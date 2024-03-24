package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(length of list1 + length of list2) ---> O(n)
// Space: O(length of list1 + length of list2) ---> O(n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}
