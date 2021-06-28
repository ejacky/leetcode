package _02_binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var tree *TreeNode
	queue = append(queue, root)

	var level [][]int
	var record []int

	for len(queue) > 0 {
		tree = queue[0]
		queue = queue[1:]

		record = append(record, tree.Val)

		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}

		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}

		level = append(level, record)

	}

	return level

}

func leverInsert(nums []int) *TreeNode {
	var root *TreeNode
	root = &TreeNode{Val: nums[0]}

	var treeQueue []*TreeNode
	treeQueue = append(treeQueue, root)
	for i := 1; i <= len(nums)-1; i++ {
		tree := treeQueue[0]
		treeQueue = treeQueue[1:]
		if nums[i] != -1 {
			tree.Left = &TreeNode{Val: nums[i]}
			treeQueue = append(treeQueue, tree.Left)
		}
		if i+1 > len(nums)-1 {
			break
		}

		if nums[i+1] != -1 {
			tree.Right = &TreeNode{Val: nums[i+1]}
			treeQueue = append(treeQueue, tree.Right)
		}
		i++
	}

	return root
}
