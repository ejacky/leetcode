package _11_minimum_depth_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}

func minDepth_tmp(root *TreeNode) int {
	var finished bool
	var level func(root *TreeNode, depth int)
	var hasLeft, hasRight bool

	level = func(root *TreeNode, depth int) {
		if root != nil && !finished {

			if root.Left != nil {
				hasLeft = true

			}
			level(root.Left, depth+1)
			if root.Right != nil {
				hasRight = true

			}
			level(root.Right, depth+1)
		} else if root == nil && hasLeft && hasRight {
			finished = true
		}

		return
	}

	level(root, 0)
	return 1

}

func dfs(root *TreeNode, index int, res *[][]int) {
	if root == nil {
		return
	}
	if index == len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[index] = append((*res)[index], root.Val)
	dfs(root.Left, index+1, res)
	dfs(root.Right, index+1, res)
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
