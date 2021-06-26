package _32_implement_queue_using_stacks

type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var value int
	if len(this.stack) > 0 {
		value = this.stack[0]
		this.stack = this.stack[1:]
	} else {
		value = 0
	}
	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stack) > 0 {
		return this.stack[0]
	} else {
		return 0
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
