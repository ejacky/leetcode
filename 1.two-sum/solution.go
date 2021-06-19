package two_sum

func twoSum(nums []int, target int) []int {
	for index := 0; index < len(nums)-1; index++ {
		for index2 := index + 1; index2 < len(nums); index2++ {
			if (nums[index] + nums[index2]) == target {
				return []int{index, index2}
			}
		}
	}
	return []int{-1, -1}

}

func twoSum1(nums []int, target int) []int {

	lens := len(nums)
	for i, j := 0, lens-1; i < j; i, j = i+1, j-1 {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
	}
	return []int{-1, -1}

}
