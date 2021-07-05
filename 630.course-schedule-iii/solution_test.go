package _07_course_schedule

import "fmt"

func ExampleScheduleCourse() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
	//output:
	//3
}
