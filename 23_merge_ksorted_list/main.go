package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeInfo struct {
	Node *ListNode
	Idx  int
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var (
		curr *ListNode
		head *ListNode
	)

	for {
		nodes := make([]*NodeInfo, 0)
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				nodes = append(nodes, &NodeInfo{Node: lists[i], Idx: i})
			}
		}

		if len(nodes) == 0 {
			break
		}

		minNode := findMinNode(nodes)

		if curr == nil {
			curr = minNode.Node
			head = minNode.Node
		} else {
			curr.Next = minNode.Node
			head.Next = minNode.Node
		}

		lists[minNode.Idx] = lists[minNode.Idx].Next
	}

	return head
}

func findMinNode(nodes []*NodeInfo) *NodeInfo {
	minNode := nodes[0]
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Node.Val < minNode.Node.Val {
			minNode = nodes[i]
		}
	}
	return minNode
}
