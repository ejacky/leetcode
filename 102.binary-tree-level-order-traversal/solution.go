package _02_binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 关键是要想到使用两个队列
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue1, queue2 []*TreeNode
	var tree *TreeNode
	queue1 = append(queue1, root)

	var records [][]int
	var record []int

	for len(queue1) > 0 {
		tree = queue1[0]
		queue1 = queue1[1:]

		record = append(record, tree.Val)

		if tree.Left != nil {
			queue2 = append(queue2, tree.Left)
		}

		if tree.Right != nil {
			queue2 = append(queue2, tree.Right)
		}

		if len(queue1) == 0 {
			records = append(records, record)
			record = []int{} // 不能写成 record[:0]
			if len(queue2) > 0 {
				queue1 = append(queue1, queue2...)
				queue2 = queue2[:0]
			}
		}
	}

	return records

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
