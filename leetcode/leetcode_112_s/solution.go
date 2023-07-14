package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := make([]*TreeNode, 1)
	visited := make([]*TreeNode, 0)
	stack[0] = root

	sum := 0

	for len(stack) != 0 {
		node := stack[len(stack)-1]

		if node.Left == nil && node.Right == nil { // 叶子节点
			sum += node.Val
			if sum == targetSum {
				return true
			}
			stack = stack[:len(stack)-1]
			sum -= node.Val
		} else { // 非叶子节点
			if len(visited) == 0 || node != visited[len(visited)-1] { // 没有向下遍历过
				sum += node.Val
				if node.Right != nil {
					stack = append(stack, node.Right)
				}

				if node.Left != nil {
					stack = append(stack, node.Left)
				}

				visited = append(visited, node)
			} else if node == visited[len(visited)-1] { // 已经向下遍历过
				stack = stack[:len(stack)-1]
				visited = visited[:len(visited)-1]
				sum -= node.Val
			}
		}
	}

	return false
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val:   -2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(hasPathSum(root, 5))
}
