package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

var minDiff int = math.MaxInt64
var prev *TreeNode = nil

func minDifferenct(root *TreeNode) {
	if root == nil {
		return
	}

	minDifferenct(root.Left)

	if prev != nil {
		if diff := abs(prev.Val - root.Val); diff < minDiff {
			minDiff = diff
		}
	}

	prev = root
	minDifferenct(root.Right)
}

func getMinimumDifference(root *TreeNode) int {
	minDifferenct(root)
	diff := minDiff
	minDiff = math.MaxInt64
	prev = nil
	return diff
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   49,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(getMinimumDifference(root))
}
