package main

type Node struct {
	Val  int
	Key  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	MaxCap  int
	CurrCap int
	Cache   map[int]*Node
	Tail    *Node
	Head    *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		MaxCap:  capacity,
		CurrCap: 0,
		Cache:   make(map[int]*Node),
		Tail:    nil,
		Head:    nil,
	}
}

// If key exists in cache, remove target node from linked list.
// prepend target to to the front of the linked list
// update key:node cache map.
func (this *LRUCache) Get(key int) int {
	tNode, exists := this.Cache[key]

	if !exists {
		return -1
	}

	this.removeNode(tNode)

	if this.Head == nil {
		this.Head, this.Tail = tNode, tNode
	} else {
		tNode.Next = this.Head
		this.Head.Prev = tNode.Next
		this.Head = tNode
	}

	this.Cache[key] = tNode

	return tNode.Val
}

func (this *LRUCache) removeNode(node *Node) {
	prevNode := node.Prev
	nextNode := node.Next

	if prevNode != nil {
		prevNode.Next = nextNode
	} else {
		// 要是我前個 node 不存在，那我們就是刪除 head node
		this.Head = nextNode
	}

	if nextNode != nil {
		nextNode.Prev = prevNode
	} else {
		// 要是我下個 node 不存在，那我們就是刪除 tail node
		this.Tail = prevNode
	}

	// remove node binding
	node.Next = nil
	node.Prev = nil
}

// If key exists in cache
//   - remove existing node from linked list
//   - prepend new node to head of linked list
//   - update cache
//
// If key not exists in cache
//
//	if capacity less than max, increment curr cap by 1, create new node, prepend new node to linked list, update key:node cache
//	if capacity is >= max, increment remove tail node including cache. create new node, prepend new node to linked list, update key:node cache
func (this *LRUCache) Put(key int, value int) {
	tNode, exists := this.Cache[key]
	if exists {
		this.removeNode(tNode)
	} else {
		if this.CurrCap < this.MaxCap {
			this.CurrCap += 1
		} else {
			tail := this.Tail
			this.removeNode(tail)
			delete(this.Cache, tail.Key)
		}
	}

	newNode := &Node{Key: key, Val: value}

	if this.Head == nil {
		this.Head, this.Tail = newNode, newNode
	} else {
		newNode.Next = this.Head
		this.Head.Prev = newNode
		this.Head = newNode
	}

	this.Cache[key] = newNode

}
