package main

import (
	"log"
	"testing"
)

// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

func TestLRUCache2(t *testing.T) {
	lru := Constructor2(2)
	// (1) -> (2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	log.Printf("debug 1 %v", lru.Cache)
	log.Printf("debug 2 %v %v", lru.Head, lru.Head.Val)
	log.Printf("debug 3 %v %v", lru.Tail, lru.Tail.Val)

	lru.Get(1)

	// log.Printf("debug 1-1 %v", lru.Cache)
	// log.Printf("debug 2-1 %v %v", lru.Head, lru.Head.Val)
	// log.Printf("debug 3-1 %v %v", lru.Tail, lru.Tail.Val)

	lru.Put(3, 3)
	log.Printf("debug 1-2 %v", lru.Cache)
	log.Printf("debug 2-2 %v %v", lru.Head, lru.Head.Val)
	log.Printf("debug 3-2 %v %v", lru.Tail, lru.Tail.Val)
}
