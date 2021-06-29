package _38_range_sum_of_bst

import (
	"fmt"
)

func ExampleRangeSumBST() {
	fmt.Println(rangeSumBST(createBstTree([]int{10, 5, 15, 3, 7, 18}), 7, 15))
	fmt.Println(rangeSumBST(createBstTree([]int{10, 5, 15, 3, 7, 13, 18, 1, 6}), 6, 10))

	//output:
	//32
	//23

}
