package main

func copyRandomList2(head *Node) *Node {
	dummy := &Node{}
	newHead := dummy
	origHead := head

	// map of orig node -> new node
	nodeMap := map[*Node]*Node{}

	for origHead != nil {
		newNode := &Node{Val: origHead.Val}
		newHead.Next = newNode

		nodeMap[origHead] = newNode
		newHead = newHead.Next
		origHead = origHead.Next
	}

	newHead2 := dummy.Next
	origHead2 := head

	for origHead2 != nil {
		if origHead2.Random != nil {
			randNewNode := nodeMap[origHead2.Random]
			newHead2.Random = randNewNode

		}

		origHead2 = origHead2.Next
		newHead2 = newHead2.Next
	}

	return dummy.Next
}
