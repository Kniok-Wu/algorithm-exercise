package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{ // 后序结尾一定是根节点
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}

	left := 0

	for inorder[left] != root.Val {
		left += 1
	}

	if len(inorder) != 1 {
		root.Left = buildTree(inorder[:left], postorder[:left])

		root.Right = buildTree(inorder[left+1:], postorder[left:len(postorder)-1])
	}

	return root
}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func main() {
	inorder := []int{6, 4, 7, 2, 8, 5, 9, 1, 12, 10, 3, 11}
	postorder := []int{6, 7, 4, 8, 9, 5, 2, 12, 10, 11, 3, 1}
	res := buildTree(inorder, postorder)
	preorderTraversal(res)
}
