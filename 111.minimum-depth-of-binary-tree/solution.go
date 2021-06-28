package _11_minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var depth int
	var invert func(root *TreeNode)

	invert = func(root *TreeNode) {
		if root != nil {
			invert(root.Left)
			invert(root.Right)
		}
	}

	invert(root)
	return depth

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
