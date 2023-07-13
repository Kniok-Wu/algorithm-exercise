package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 1)
	result[0] = make([]int, 0)

	queue := make([]*TreeNode, 1)
	queue[0] = root

	layerNum := 1

	var node *TreeNode = nil

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		result[len(result)-1] = append(result[len(result)-1], node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		layerNum -= 1

		if layerNum == 0 && len(queue) != 0 {
			layerNum = len(queue)
			result = append(result, make([]int, 0))
		}
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(levelOrder(root))
}
