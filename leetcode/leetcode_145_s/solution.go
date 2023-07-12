package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	visited := make([]*TreeNode, 0)
	res := make([]int, 0)

	for cur := root; cur != nil; cur = cur.Left {
		stack = append(stack, cur)
	} // 添加根节点下所有的左节点

	for len(stack) != 0 {
		node := stack[len(stack)-1]

		if len(visited) == 0 || node != visited[len(visited)-1] { //如果没有遍历过该节点的右节点 则尝试遍历右节点的所有左节点
			if node.Right != nil {
				visited = append(visited, node)
				node = node.Right
				for node != nil { // 遍历该右节点的所有左节点
					stack = append(stack, node)
					node = node.Left
				}
			} else { // 没有右节点 直接输出
				res = append(res, node.Val)
				stack = stack[:len(stack)-1] // 弹出该节点
			}
		} else { // 已经遍历过该节点的右节点
			visited = visited[:len(visited)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1] // 弹出该节点

		}
	}

	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(postorderTraversal(root))
}
