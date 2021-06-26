package _25_implement_stack_using_queues

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var value int
	if len(this.queue) > 0 {
		value = this.queue[len(this.queue)-1]
		this.queue = this.queue[:len(this.queue)-1]
		return value
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
