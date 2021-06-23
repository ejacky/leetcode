package _38_range_sum_of_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var num int

	var search func(tree *TreeNode, low int, high int)
	search = func(tree *TreeNode, low int, high int) {
		if tree.Val >= low && tree.Val <= high {
			num += tree.Val
			if tree.Left != nil {
				search(tree.Left, low, high)
			}

			if tree.Right != nil {
				search(tree.Right, low, high)
			}
		}
		if tree.Val < low {
			if tree.Right != nil {
				search(tree.Right, low, high)
			}
		}
		if tree.Val > high {
			if tree.Left != nil {
				search(tree.Left, low, high)
			}
		}
	}

	search(root, low, high)

	return num

}

func rangeSumBST1(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBST2(root *TreeNode, low, high int) (sum int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		if node.Val > high {
			q = append(q, node.Left)
		} else if node.Val < low {
			q = append(q, node.Right)
		} else {
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}
	}
	return
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
	for _, num := range nums[1:] {
		insert(root, num)
	}

	return root
}
