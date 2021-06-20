package _07_course_schedule

import "fmt"

func ExampleCanFinished() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	//output:
	//true
	//false
}
