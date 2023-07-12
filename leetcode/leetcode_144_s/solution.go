package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	result := make([]int, 0)
//	result = append(result, root.Val)
//	resL := preorderTraversal(root.Left)
//	resR := preorderTraversal(root.Right)
//	if resL != nil {
//		result = append(result, resL...)
//	}
//
//	if resR != nil {
//		result = append(result, resR...)
//	}
//
//	return result
//}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	res := make([]int, 0)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

func main() {
	fmt.Println([]int{1}[1:])
}
