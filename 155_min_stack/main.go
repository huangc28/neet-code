package main

/*
push:

	append element at the end of the stack

pop:

	pop the last element in the stack if stack length is > 0

Top:

	return

[-2,0,-3,-1,~-100,~-99,~-98],
push, push, push, push, push , push, push
pop, pop, pop, getMin, pop

curr_min_pos:~4,2

pos:prev_min_pos

	{
	    0:0
	    1:0
	    2:0
	    3:2 // min ---> stack[2], -3
	    ~4:2
	    ~5:4
	    ~6:4
	}

- we'll solve it using a stack that stores the value and hash map of pos:prev_min_pos
- when push a new value, we check if the new value is greater than current min val
  - greater, don't change current min pos, set minMap {curr_pos:curr_min_pos}
  - smaller, change current min val to be current positin, set minMap {curr_pos:prev_min_pos}

- When pop a new value, we retrieve the previous min position from prevMinPosMap by `prevMinPosMap[len(stack)-1]`
*/
type MinStack struct {
	// Stores the position of current min value in the stack.
	currMinPos int

	// Stores the values pushed
	stack []int

	// Store the previous min position for each position in the stack.
	// For example, when we pop the last position element from the stack,
	// we'll replace the currMinPos to be preMinPosMap[len(stack)-1]
	prevMinPosMap map[int]int
}

func Constructor() MinStack {
	return MinStack{
		stack:         make([]int, 0),
		prevMinPosMap: make(map[int]int),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val)
		this.currMinPos = 0
		this.prevMinPosMap[0] = 0

		return
	}

	this.stack = append(this.stack, val)
	this.prevMinPosMap[len(this.stack)-1] = this.currMinPos

	if val < this.stack[this.currMinPos] {
		this.currMinPos = len(this.stack) - 1
	}
}
func (this *MinStack) Pop() {
	lastPos := len(this.stack) - 1
	this.stack = this.stack[:lastPos]

	this.currMinPos = this.prevMinPosMap[lastPos]
	delete(this.prevMinPosMap, lastPos)
}
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
	return this.stack[this.currMinPos]
}
