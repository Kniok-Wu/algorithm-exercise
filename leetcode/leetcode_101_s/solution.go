package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	result := make([]int, 0)

	queue := make([]*TreeNode, 1)
	queue[0] = root

	layerNum := 1

	var node *TreeNode = nil

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		layerNum -= 1
		if node == nil {
			result = append(result, -101)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}

		if layerNum == 0 { // 该层遍历结束 且存在下层
			layerNum = len(queue)
			for i, j := 0, len(result)-1; i <= j; {
				if result[i] != result[j] {
					return false
				}
				i += 1
				j -= 1
			}
			result = make([]int, 0)
		}
	}

	return true
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Println(isSymmetric(root))
}
