package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList2(head *ListNode) {
	listLen := calcLen(head)
	half := listLen / 2

	var list1 *ListNode
	curr := head
	for i := 0; i < half; i++ {
		if list1 == nil {
			list1 = curr
		} else {
			list1.Next = curr
			list1 = list1.Next
		}
		curr = curr.Next
	}

	list2 := list1.Next
	list1.Next = nil

	list2 = reverse(list2)
	head = mergeList(head, list2)

	// for head != nil {
	// 	log.Printf("debug ~~ %v", head.Val)
	// 	head = head.Next
	// }
	//
	// for list2 != nil {
	// 	log.Printf("debug ** %v", list2.Val)
	// 	list2 = list2.Next
	// }

	// for hh != nil {
	// 	log.Printf("debug 1 %v", hh)
	// 	hh = hh.Next
	// }
	// len1 := calcLen(head)
	// len2 := calcLen(list2)

	// log.Printf("debug 2 %v ", len2)
}

func calcLen(head *ListNode) int {
	var l int
	curr := head

	for curr != nil {
		l += 1
		curr = curr.Next
	}

	return l
}

func mergeList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	log.Printf("debug 1 %v", calcLen(list1))
	log.Printf("debug 2 %v", calcLen(list2))

	for list1 != nil && list2 != nil {
		curr.Next = list1
		list1 = list1.Next
		curr = curr.Next

		curr.Next = list2
		curr = curr.Next
		list2 = list2.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	log.Printf("debug 4 %v", calcLen(list2))
	log.Printf("debug 1 %v", calcLen(list1))

	return dummy.Next
}
