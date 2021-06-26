package _00_same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return false
}

func createTree(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) == 0 {
		return root
	}

	var insert func(tree *TreeNode, num int)
	insert = func(tree *TreeNode, num int) {

		if num < tree.Val {
			if tree.Left == nil {
				tree.Left = &TreeNode{Val: num}
			} else {
				insert(tree.Left, num)
			}
		}

		if num > tree.Val {
			if tree.Right == nil {
				tree.Right = &TreeNode{Val: num}
			} else {
				insert(tree.Right, num)
			}
		}
	}

	root = &TreeNode{}
	root.Val = nums[0]

	var queueNode []TreeNode

	for i := 0; i < len(nums); i++ {
		queueNode = append(queueNode, TreeNode{Val: nums[i]})
	}

	for _, num := range nums[1:] {
		insert(root, num)
	}

	return root
}
