package add_two_numbers

import "fmt"

//func ExampleAddTwoNumbers() {
//	nums := []int{2, 4, 3}
//	l1 := createRearList(nums)
//	nums1 := []int{5, 6, 4}
//	l2 := createRearList(nums1)
//
//	fmt.Println(addTwoNumbers(l1, l2))
//	//output:
//	//[7 0 8]
//
//}

func ExampleAddTwoNumbers() {
	nums := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 := createRearList(nums)
	nums1 := []int{5, 6, 4}
	l2 := createRearList(nums1)

	fmt.Println(lSlice(addTwoNumbers(l1, l2)))
	//output:
	//[6 6 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

}
