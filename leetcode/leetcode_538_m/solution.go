package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func accumulate(root *TreeNode, parent int) int {
	if root == nil {
		return parent
	}

	root.Val += accumulate(root.Right, parent)

	return accumulate(root.Left, root.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	accumulate(root, 0)
	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	convertBST(root)
}
