package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if root == nil {
		return newNode
	}

	var prev *TreeNode = nil
	cur := root

	for cur != nil {
		prev = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if prev.Val > val {
		newNode.Left = prev.Left
		prev.Left = newNode
	} else {
		newNode.Right = prev.Right
		prev.Right = newNode
	}

	return root
}
