package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	node  *TreeNode
	depth int
}

func findBottomLeftValue(root *TreeNode) int {
	stack := make([]*TreeNode, 1)
	stack[0] = root

	result := &Result{
		node:  root,
		depth: 1,
	}

	layerNum := 1
	depth := 2

	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node.Left != nil {
			stack = append(stack, node.Left)
			if result.depth < depth {
				result.depth = depth
				result.node = node.Left
			}
		}
		if node.Right != nil {
			stack = append(stack, node.Right)

			if result.depth < depth {
				result.depth = depth
				result.node = node.Right
			}
		}
		layerNum -= 1
		if layerNum == 0 {
			layerNum = len(stack)
			depth += 1
		}
	}
	return result.node.Val
}

func main() {
	root := &TreeNode{
		Val:  0,
		Left: nil,
		Right: &TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(findBottomLeftValue(root))
}
