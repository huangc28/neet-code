package main

/*
3 -> 4
h

capacity: 2
size: ~0,~1,~2,1,2
get:
{
    ~1: 1
    ~2: 2
    3: 3
    4: 4
}

output:
    get(1): 1
    get(2): 2

- When get is performed, move the node to the tail of the list
- When put is performed
    - when size is < capacity, append the node to the tail of the list
    - when size is >= capacity, remove the head of the linked list, append the node to the tail of the list
*/

type Node2 struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache2 struct {
	capacity int
	size     int
	Head     *Node
	Tail     *Node
	Cache    map[int]*Node
}

func Constructor2(capacity int) LRUCache2 {
	return LRUCache2{
		capacity: capacity,
		size:     0,
		Head:     nil,
		Tail:     nil,
		Cache:    make(map[int]*Node),
	}
}

func (this *LRUCache2) Get(key int) int {
	// if key not exists in cache, return -1
	node, exists := this.Cache[key]

	if !exists {
		return -1
	}

	// if key exists in cache, we need to move
	// the corresponding node to the tail of the list
	// 1. delete node in the link list
	//   1 <-> 2 <-> 3, remove 2
	//   1 <-> 3
	// 2. append node to the end of the list
	//    1 <-> 3 <-> 2
	this.delete(node)

	this.append(node)

	return node.Val
}

// Delete node in the doubly link list
//
//	1 <-> 2 <-> 3, remove 2
//	1 <-> 3
func (this *LRUCache2) delete(node *Node) {
	// if node's previous node is not nil, point prev node next to curr node's next
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// the node is head, since it does not have prev. change the head to node's next
		this.Head = node.Next
	}

	// if node's next node is not nil, point next node's prev to curr node's prev
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// the node is tail since it does not have next. change the tail to node's prev
		this.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}

// Append node to the end of the list
//
//	1 <-> 3 <-> 2
func (this *LRUCache2) append(node *Node) {
	// No linked list exists yet. both head and tail is the node
	if this.Head == nil {
		this.Head, this.Tail = node, node
	} else {
		// linked list exists, append node to the tail of the list
		this.Tail.Next = node
		node.Prev = this.Tail
		this.Tail = node
	}
}

func (this *LRUCache2) Put(key int, value int) {
	// if key already exists in cache, update cache, place node in tail of linked list
	node, exists := this.Cache[key]

	if exists {
		node.Val = value

		this.delete(node)
		this.append(node)
	} else {
		node := &Node{Key: key, Val: value}

		// if curr size is less than capacity:
		// 1. append node to end of list
		// 2. add key to cache
		if this.size < this.capacity {
			this.append(node)
			this.Cache[key] = node
			this.size += 1
		} else {
			// if curr size is >= capacity
			// 1. locate the head node in the cache, remove it.
			// 2. evict head node in linked list
			// 3. append new node to the linked list
			delete(this.Cache, this.Head.Key)

			this.delete(this.Head)

			this.append(node)

			this.Cache[key] = node
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
