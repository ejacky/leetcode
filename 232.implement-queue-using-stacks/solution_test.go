package _32_implement_queue_using_stacks

import "fmt"

func Example232() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
	//output:
	//1
	//1
	//false
}
