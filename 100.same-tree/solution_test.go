package _00_same_tree

import "fmt"

func ExampleSameTree() {
	fmt.Println(isSameTree(leverInsert([]int{1, 2, 3}), leverInsert([]int{1, 2, 3})))
	fmt.Println(isSameTree(leverInsert([]int{1, 2}), leverInsert([]int{1, -1, 2})))
	fmt.Println(isSameTree(leverInsert([]int{1, 2, 1}), leverInsert([]int{1, 1, 2})))
	//output:
	//true
	//false
	//false
}
