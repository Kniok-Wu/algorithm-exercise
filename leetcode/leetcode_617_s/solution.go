package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	switch {
	case root1 == nil && root2 == nil:
		return nil
	case root1 != nil && root2 == nil:
		return root1
	case root1 == nil && root2 != nil:
		return root2
	default:
		return &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
	}
}

func main() {

}
