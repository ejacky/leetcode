package _25_implement_stack_using_queues

import "fmt"

func Example225() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
	//output:
	//2
	//2
	//false
}
