package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		} else {
			res += sumOfLeftLeaves(root.Left)
		}
	}
	res += sumOfLeftLeaves(root.Right)

	return res
}
