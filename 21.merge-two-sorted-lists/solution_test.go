package _1_merge_two_sorted_lists

import "fmt"

func ExampleMergeTwoLists() {
	var nums []int
	nums = []int{1, 2, 4}
	l1 := createRearList(nums)
	nums = []int{1, 3, 4}
	l2 := createRearList(nums)
	fmt.Println(lSlice(mergeTwoLists(l1, l2)))

	//output:
	//[1 1 2 3 4 4]
}
