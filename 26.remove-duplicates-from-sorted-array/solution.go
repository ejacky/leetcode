package _6_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var flag bool
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			flag = true
			continue
		} else {
			if flag {
				j++
				nums[j] = nums[i]
			} else {
				j++
			}
		}
	}
	return j + 1
}
