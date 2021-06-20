package _10_course_schedule_ii

import "fmt"

func ExampleFindOrder() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(5, [][]int{}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(3, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{3, 0}, {0, 1}}))

	//output:
	//[0 1]
	//[4 3 2 1 0]
	//[0 1 2 3]
	//[2 0 1]
	//[2 1 0 3]
}
