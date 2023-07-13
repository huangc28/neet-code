package main

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
1: [2,4]
2: [1,3]
*/

func cloneGraph(node *Node) *Node {
	adjList := make(map[int][]int)
	copiedNodes := make(map[int]*Node)

	stack := make([]*Node, 0)
	stack = append(stack, node)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, exists := adjList[node.Val]; exists {
			continue
		}

		adjList[node.Val] = make([]int, 0)
		copiedNodes[node.Val] = &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, 0),
		}

		for _, neighbor := range node.Neighbors {
			adjList[node.Val] = append(adjList[node.Val], neighbor.Val)
			stack = append(stack, neighbor)
		}
	}

	for val, neighborVals := range adjList {
		copiedNode := copiedNodes[val]
		for _, neighborVal := range neighborVals {
			copiedNode.Neighbors = append(
				copiedNode.Neighbors,
				copiedNodes[neighborVal],
			)
		}
	}

	return copiedNodes[node.Val]
}
