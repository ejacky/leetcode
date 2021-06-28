package _02_binary_tree_level_order_traversal

import "fmt"

func ExampleLevelOrder() {
	fmt.Println(levelOrder(leverInsert([]int{3, 9, 20, -1, -1, 15, 7})))
	//output:
	//[[3] [9 20] [15 7]]
}
