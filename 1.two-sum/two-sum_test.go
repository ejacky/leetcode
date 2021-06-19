package two_sum

import "fmt"

func ExampleTwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 5}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	//output:
	//[0 1]
	//[1 2]
	//[0 1]
}

func ExampleTwoSum1() {
	fmt.Println(twoSum1([]int{2, 7, 11, 5}, 9))
	fmt.Println(twoSum1([]int{3, 2, 4}, 6))
	fmt.Println(twoSum1([]int{3, 3}, 6))
	//output:
	//[0 1]
	//[1 2]
	//[0 1]
}
