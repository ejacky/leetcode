package _37_delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

var nums = []int{4, 5, 1, 9}
var head = createRearList(nums)

func deleteNode(node *ListNode) {

	var rear = head
	var tmp *ListNode

	for rear != nil {
		if rear.Val == node.Val {
			tmp.Next = rear.Next
			break
		}
		tmp = rear

		rear = rear.Next
	}

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

func lSlice(l *ListNode) []int {
	var lslice []int
	for l != nil {
		lslice = append(lslice, l.Val)
		l = l.Next
	}
	return lslice
}
