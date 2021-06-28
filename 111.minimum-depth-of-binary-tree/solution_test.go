package _11_minimum_depth_of_binary_tree

import "fmt"

func ExampleMinDepth() {
	fmt.Println(minDepth(leverInsert([]int{3, 9, 20, -1, -1, 15, 7})))
	fmt.Println(minDepth(leverInsert([]int{2, -1, 3, -1, 4, -1, 5, -1, 6})))
	//output:
	//2
	//5
}
