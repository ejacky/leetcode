package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1 []int

	for i := 0; l1 != nil; i++ {
		num1 = append(num1, l1.Val)
		l1 = l1.Next
	}

	var num2 []int
	for i := 0; l2 != nil; i++ {
		num2 = append(num2, l2.Val)
		l2 = l2.Next
	}

	var nums []int
	var t1, t2 int
	var flag bool
	for i, j := 0, 0; !(i > len(num1)-1 && j > len(num2)-1); {
		if i <= len(num1)-1 {
			t1 = num1[i]
			i++
		} else {
			t1 = 0
		}
		if j <= len(num2)-1 {
			t2 = num2[j]
			j++
		} else {
			t2 = 0
		}

		if flag {
			if t1+t2+1 >= 10 {
				flag = true
				nums = append(nums, t1+t2+1-10)
			} else {
				flag = false
				nums = append(nums, t1+t2+1)
			}
		} else {
			if t1+t2 >= 10 {
				flag = true
				nums = append(nums, t1+t2-10)
			} else {
				flag = false
				nums = append(nums, t1+t2)
			}
		}
	}

	if flag {
		nums = append(nums, 1)
	}

	return createRearList(nums)
}

func createRearList(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}

	var l, r *ListNode
	l = &ListNode{}
	r = l
	r.Val = nums[0]
	for _, num := range nums[1:] {
		node := &ListNode{}
		node.Val = num
		r.Next = node
		r = node
	}
	r.Next = nil
	return l
}

func createFrontList(nums []int) *ListNode {
	var l *ListNode
	l = &ListNode{}
	l.Next = nil
	for _, num := range nums {
		node := &ListNode{}
		node.Val = num
		node.Next = l.Next
		l.Next = node
	}
	return l
}
