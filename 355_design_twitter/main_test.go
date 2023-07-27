package main

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	twt := Constructor()
	twt.PostTweet(1, 5)
	res1 := twt.GetNewsFeed(1)
	log.Printf("debug res1 %v", res1)
	twt.Follow(1, 2)
	twt.PostTweet(2, 6)
	res2 := twt.GetNewsFeed(1)
	log.Printf("debug res2 %v", res2)
}

func TestCase2(t *testing.T) {
	twt := Constructor()
	twt.PostTweet(1, 5)
	twt.PostTweet(1, 3)
	res1 := twt.GetNewsFeed(1)
	log.Printf("debug 1 %v", res1)
}
