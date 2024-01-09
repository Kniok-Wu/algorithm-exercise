/**
 * @Time: 2024/1/8 17:35
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	stack := make([]*TreeNode, 1)
	stack[0] = root

	for i := 0; i < len(stack); i++ {
		if stack[i].Left != nil {
			stack = append(stack, stack[i].Left)
		}

		if stack[i].Right != nil {
			stack = append(stack, stack[i].Right)
		}
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left == nil && node.Right == nil {
			continue
		}

		max := node.Val

		if node.Left == nil {
			node.Left = &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}
		}

		if node.Right == nil {
			node.Right = &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}
		}

		// 1. 选择直系子节点
		if node.Left.Val+node.Right.Val > max {
			max = node.Left.Val + node.Right.Val
		}

		// 2. 选择孙节点
		grandsonVal := node.Val
		if node.Left.Left != nil {
			grandsonVal += node.Left.Left.Val
		}
		if node.Left.Right != nil {
			grandsonVal += node.Left.Right.Val
		}
		if node.Right.Left != nil {
			grandsonVal += node.Right.Left.Val
		}
		if node.Right.Right != nil {
			grandsonVal += node.Right.Right.Val
		}

		if grandsonVal > max {
			max = grandsonVal
		}

		node.Val = max
	}

	return root.Val
}

func robII(root *TreeNode) int {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		dp := make([]int, 2)
		left := make([]int, 0, 2)
		right := make([]int, 0, 2)

		if node.Left != nil {
			left = append(left, dfs(node.Left)...)
		} else {
			left = append(left, []int{0, 0}...)
		}

		if node.Right != nil {
			right = append(right, dfs(node.Right)...)
		} else {
			right = append(right, []int{0, 0}...)
		}

		// 1. 0 表示不偷当前节点
		var maxLeft, maxRight int
		if left[0] > left[1] {
			maxLeft = left[0]
		} else {
			maxLeft = left[1]
		}

		if right[0] > right[1] {
			maxRight = right[0]
		} else {
			maxRight = right[1]
		}

		dp[0] = maxLeft + maxRight

		// 1. 表示偷当前节点
		dp[1] = node.Val + left[0] + right[0]

		return dp
	}

	dp := dfs(root)

	if dp[0] > dp[1] {
		return dp[0]
	} else {
		return dp[1]
	}
}

func main() {
	fmt.Println(robII(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}))
}
