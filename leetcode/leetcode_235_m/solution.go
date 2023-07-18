package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ancestor *TreeNode = nil

func Traversal(root, p, q *TreeNode) (bool, bool) {
	findP, findQ := false, false

	switch {
	case root == nil: // 空值没有任何操作
	case root.Val == p.Val: // 当前节点值等于小值 只查找右子树
		findP = true
		_, findQ = Traversal(root.Right, p, q)
	case root.Val == q.Val: // 当前节点值等于大值 只查找左子树
		findQ = true
		findP, _ = Traversal(root.Left, p, q)
	case root.Val < p.Val: // 当前节点比最小值还要小 只查找右子树
		findP, findQ = Traversal(root.Right, p, q)
	case root.Val > q.Val: // 当前节点比最大值还要大 只查找左子树
		findP, findQ = Traversal(root.Right, p, q)
	default: // 当前节点在两值之间 左右子树均查找
		findP, findQ = Traversal(root.Left, p, q) // 遍历左子树
		resP, resQ := Traversal(root.Right, p, q)
		findP, findQ = resP || findP, resQ || findQ
	}

	if findP && findQ { // 如果均找到了
		if ancestor == nil {
			ancestor = root
		}
	}

	return findP, findQ
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val { // 保证节点p的值较小
		p.Val, q.Val = q.Val, p.Val
	}

	Traversal(root, p, q)

	res := ancestor
	ancestor = nil
	return res
}
