package _1_merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var hNode, rNode *ListNode
	rNode = &ListNode{}
	hNode = rNode

	if l1 == nil && l2 == nil {
		return nil
	}
	for !(l1 == nil && l2 == nil) {
		var node ListNode
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				node.Val = l1.Val
				l1 = l1.Next
			} else {
				node.Val = l2.Val
				l2 = l2.Next
			}

		} else if l1 == nil && l2 != nil {
			node.Val = l2.Val
			l2 = l2.Next

		} else if l1 != nil && l2 == nil {
			node.Val = l1.Val
			l1 = l1.Next
		}

		rNode.Next = &node
		rNode = &node
	}
	return hNode.Next

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
