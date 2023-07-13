package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else {
		lDepth := minDepth(root.Left)
		rDepth := minDepth(root.Right)
		if lDepth > rDepth {
			return rDepth + 1
		} else {
			return lDepth + 1
		}
	}
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(minDepth(root))
}
