package main

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	for len(lists) > 1 {
		list1, list2 := lists[0], lists[1]
		mergedList := mergeLists(list1, list2)
		lists = lists[2:]
		lists = append(lists, mergedList)
	}

	return lists[0]
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	head := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}

	if list2 != nil {
		head.Next = list2
	}

	return dummy.Next
}
