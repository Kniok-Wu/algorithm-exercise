package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for cur := root; cur != nil; {
		if cur.Val == val {
			return cur
		} else if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return nil
}
