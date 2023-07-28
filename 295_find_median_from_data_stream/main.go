package main

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

type MaxHeap struct {
	Nums []int
}

func (h *MaxHeap) Top() int {
	return h.Nums[0]
}

func (h *MaxHeap) Add(num int) {
	h.Nums = append(h.Nums, num)
	h.siftUp()
}

func (h *MaxHeap) siftUp() {
	curr := len(h.Nums) - 1
	for curr >= 0 {
		parent := (curr - 1) / 2
		if h.Nums[curr] > h.Nums[parent] {
			swap(h.Nums, curr, parent)
			curr = parent
		} else {
			return
		}
	}
}

func (*MaxHeap) Pop() int {
	return 0
}

type MinHeap struct {
	Nums []int
}

func (h *MinHeap) Top() int {
	return h.Nums[0]
}

func (h *MinHeap) Add(num int) {
}

func (*MinHeap) Pop() int {
	return 0
}

type MedianFinder struct {
	smaller *MaxHeap // max heap
	greater *MinHeap // min heap
	median  float64
}

func Constructor() MedianFinder {
	return MedianFinder{
		smaller: &MaxHeap{Nums: make([]int, 0)},
		greater: &MinHeap{Nums: make([]int, 0)},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.smaller.Nums) == 0 && len(this.greater.Nums) == 0 {
		this.median = float64(num)
	}

	// add number to smaller or greater
	if num > this.smaller.Top() {
		this.greater.Add(num)
	} else {
		this.smaller.Add(num)
	}

	if len(this.greater.Nums) == len(this.smaller.Nums) {
		this.median = (float64(this.greater.Top()) + float64(this.smaller.Top())) / 2
	} else if len(this.greater.Nums) > len(this.smaller.Nums) {
		this.median = float64(this.greater.Top())
	} else {
		this.median = float64(this.smaller.Top())
	}

	// do we need to rebalance the heap
	sLen := len(this.smaller.Nums)
	gLen := len(this.greater.Nums)
	diff := sLen - gLen
	if diff > 1 {
		this.rebalance()
	}
}

func (this *MedianFinder) rebalance() {
	if len(this.greater.Nums) > len(this.smaller.Nums) {
		popped := this.greater.Pop()
		this.smaller.Add(popped)
	} else {
		popped := this.smaller.Pop()
		this.greater.Add(popped)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	return this.median
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
