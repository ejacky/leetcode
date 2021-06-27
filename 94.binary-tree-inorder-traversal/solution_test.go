package _4_binary_tree_inorder_traversal

import "fmt"

func ExampleInorderTraversal() {
	fmt.Println(inorderTraversal(leverInsert([]int{1, -1, 2, 3})))
	//output:
	//[1 3 2]
}
