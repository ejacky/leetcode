package _00_same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	var flag = true

	var preorder func(tree1 *TreeNode, tree2 *TreeNode)
	preorder = func(tree1 *TreeNode, tree2 *TreeNode) {
		if tree1 != nil && tree2 != nil {
			if tree1.Val != tree2.Val && flag {
				flag = false
			} else {
				preorder(tree1.Left, tree2.Left)
				preorder(tree1.Right, tree2.Right)
			}
		}
		if tree1 == nil && tree2 != nil || tree1 != nil && tree2 == nil {
			flag = false
		}
	}
	preorder(p, q)
	return flag
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
