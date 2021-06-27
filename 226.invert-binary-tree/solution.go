package _26_invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	var invert func(root *TreeNode)

	invert = func(root *TreeNode) {
		if root != nil {
			root.Left, root.Right = root.Right, root.Left
			invert(root.Left)
			invert(root.Right)
		}
	}

	invert(root)

	return root
}
