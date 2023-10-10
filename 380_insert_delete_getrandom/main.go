package main

type RandomizedSet struct {
	set  map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:  make(map[int]int),
		list: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {}
func (this *RandomizedSet) Remove(val int) bool {}
func (this *RandomizedSet) GetRandom() int      {}
