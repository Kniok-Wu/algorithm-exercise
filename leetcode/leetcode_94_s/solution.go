package leetcode_94_s

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for cur := root; cur != nil; cur = cur.Left {
		stack = append(stack, cur)
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Right != nil { // 如果有右节点，则将其所有左节点压栈
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}

	return res
}
