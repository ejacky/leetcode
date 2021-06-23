package _00_same_tree

import "fmt"

func ExampleSameTree() {
	fmt.Println(isSameTree(createTree([]int{1, 2, 3}), createTree([]int{1, 2, 3})))
	fmt.Println(isSameTree(createTree([]int{1, 2, 3}), createTree([]int{1, -1, 2})))
	fmt.Println(isSameTree(createTree([]int{1, 2, 1}), createTree([]int{1, 1, 2})))

	//output:
	//[0 1]
	//[1 2]
	//[0 1]
}
