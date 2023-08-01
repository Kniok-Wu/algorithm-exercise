package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{
		Val:   math.MaxInt,
		Left:  root,
		Right: nil,
	}

	cur := dummy

	for {
		if cur.Val > key { // 前往左子树
			if cur.Left == nil { // 不存在该节点
				return root
			} else if cur.Left.Val == key { // 查询左孩子节点为目标节点
				break
			} else { // 搜索左孩子树
				cur = cur.Left
			}
		} else { // 前往右子树
			if cur.Right == nil { // 不存在该节点
				return root
			} else if cur.Right.Val == key { // 查询右孩子节点为目标节点
				break
			} else { // 搜索右孩子树
				cur = cur.Right
			}
		}
	}

	var child *TreeNode
	if cur.Left != nil && cur.Left.Val == key {
		child = cur.Left.Left
		cur.Left = cur.Left.Right
	} else {
		child = cur.Right.Left
		cur.Right = cur.Right.Right
	}

	for child != nil {
		if cur.Val > child.Val { // 插入左子树
			if cur.Left == nil {
				cur.Left = child
				break
			} else {
				cur = cur.Left
			}
		} else { // 插入右子树
			if cur.Right == nil {
				cur.Right = child
				break
			} else {
				cur = cur.Right
			}
		}
	}

	return dummy.Left
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	deleteNode(root, 3)
}
