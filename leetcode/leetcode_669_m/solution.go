package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func adjust(prev, cur *TreeNode, low, high int) {
	if cur == nil {
		return
	}

	if cur.Val < low { // 当前过小 则删除该节点及其左子树
		if cur == prev.Left {
			prev.Left = cur.Right
			adjust(prev, prev.Left, low, high)
		} else {
			prev.Right = cur.Right
			adjust(prev, prev.Right, low, high)
		}
	} else if cur.Val > high { // 当前过大 则删除该节点及其右子树
		if cur == prev.Left {
			prev.Left = cur.Left
			adjust(prev, prev.Left, low, high)
		} else {
			prev.Right = cur.Left
			adjust(prev, prev.Right, low, high)
		}
	}

	if cur = prev.Left; cur != nil {
		adjust(cur, cur.Left, low, high)
		adjust(cur, cur.Right, low, high)
	}

	if cur = prev.Right; cur != nil {
		adjust(cur, cur.Left, low, high)
		adjust(cur, cur.Right, low, high)
	}
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	dummy := &TreeNode{
		Val:   math.MaxInt,
		Left:  root,
		Right: nil,
	}

	adjust(dummy, root, low, high)

	return dummy.Left
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}

	trimBST(root, 3, 4)
}
