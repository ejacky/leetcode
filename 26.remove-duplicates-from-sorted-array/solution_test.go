package _6_remove_duplicates_from_sorted_array

import "fmt"

func ExampleRemoveDuplicates() {
	var nums []int
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	lens := removeDuplicates(nums)
	fmt.Println(lens)
	fmt.Println(nums[0:lens])
	//output:
	//5
	//[0 1 2 3 4]
}
