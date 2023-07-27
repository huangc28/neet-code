package main

import "log"

type Post struct {
	id, tweetId int
}

type Twitter struct {
	// store list of followee for each follower
	// {
	//    1: {
	//	5: true // 1 follow 5
	//      4: true // 1 follow 4
	//    }
	// }
	followMap map[int]map[int]bool

	// store list of posts each user created.
	userPosts map[int][]Post

	maxPost int

	idxPointer int
}

func Constructor() Twitter {
	return Twitter{
		followMap:  make(map[int]map[int]bool),
		userPosts:  make(map[int][]Post),
		maxPost:    10,
		idxPointer: 1,
	}
}

func (this *Twitter) markFollow(userId int, follow bool) {
	if _, exists := this.followMap[userId]; !exists {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = follow
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	// add user Id to follower map
	this.markFollow(userId, true)

	// add post to userPosts map
	this.userPosts[userId] = append(this.userPosts[userId], Post{
		id:      this.idxPointer,
		tweetId: tweetId,
	})

	this.idxPointer++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followeePosts := make([]Post, 0)
	// merge all followee's posts
	// n ---> number of followers
	// m ---> nums of post
	// add -> log(n)
	// n * m * log(m)
	// n * mlog(m)
	for followeeId, followStatus := range this.followMap[userId] {
		if followStatus {
			followeePosts = append(followeePosts, this.userPosts[followeeId]...)
		}
	}

	// heapify
	MaxHeapify(followeePosts)

	for _, p := range followeePosts {
		log.Printf("debug 1 %v", p.id)
	}

	// pop 10 most recent posts
	res := make([]int, 0, 10)

	// o(n) --> can ignore
	for i := 0; i < this.maxPost; i++ {
		if len(followeePosts) > 0 {
			res = append(res, pop(&followeePosts).tweetId)
		}
	}

	return res
}

func MaxHeapify(posts []Post) {
	curr := len(posts) - 1
	for parent := (curr - 1) / 2; parent >= 0; parent-- {
		siftDown(posts, 0, len(posts)-1)
	}
}

func pop(posts *[]Post) Post {
	swap(*posts, 0, len(*posts)-1)
	n := (*posts)[len(*posts)-1]
	(*posts) = (*posts)[:len(*posts)-1]
	siftDown(*posts, 0, len(*posts)-1)
	return n
}

func siftDown(posts []Post, startIdx, endIdx int) {
	for startIdx <= endIdx {
		lc := (startIdx * 2) + 1
		rc := (startIdx * 2) + 2
		idxToSwap := -1

		if lc <= endIdx {
			idxToSwap = lc
			if rc <= endIdx && posts[rc].id > posts[lc].id {
				idxToSwap = rc
			}
		}

		if idxToSwap != -1 && posts[idxToSwap].id > posts[startIdx].id {
			swap(posts, idxToSwap, startIdx)
			startIdx = idxToSwap
		} else {
			return
		}
	}
}

func swap(posts []Post, i, j int) {
	posts[i], posts[j] = posts[j], posts[i]
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// add floweeId to fllowerId's following list
	this.markFollow(followerId, true)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	// remove floweeId from followerId's following list
	this.markFollow(followerId, false)
}
