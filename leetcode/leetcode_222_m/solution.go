package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	if root.Left == nil {
//		return 1
//	}
//
//	return 1 + countNodes(root.Left) + countNodes(root.Right)
//}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := root.Left, root.Right
	lDepth, rDepth := 0, 0

	for l != nil {
		lDepth += 1
		l = l.Left
	}

	for r != nil {
		rDepth += 1
		r = r.Right
	}

	if lDepth == rDepth {
		return 2<<lDepth - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
