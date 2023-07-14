package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Depth(root *TreeNode, balance bool) (int, bool) {
	if root == nil {
		return 0, true
	}

	if !balance {
		return -1, balance
	}

	lDepth, bfL := Depth(root.Left, balance)
	rDepth, bfR := Depth(root.Right, balance)
	bf := lDepth - rDepth

	if !(bfL && bfR) || bf <= -2 || bf >= 2 { // 如果左右孩子有任意一个不平衡 直接返回 false
		return -1, false
	}

	if lDepth > rDepth {
		return lDepth + 1, true
	} else {
		return rDepth + 1, true
	}

}

func isBalanced(root *TreeNode) bool {
	if root != nil {
		_, balance := Depth(root, true)

		return balance
	}

	return true
}

func main() {
}
