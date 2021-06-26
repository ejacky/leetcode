package _37_delete_node_in_a_linked_list

import "fmt"

func ExampleDeleteNode() {
	//
	//nums = []int{1, 3, 4}
	//l2 := createRearList(nums)

	deleteNode(&ListNode{Val: 5})
	fmt.Println(lSlice(head))

	//output:
	//[4 1 9]
}
