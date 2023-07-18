package leetcode_236_m

var lastAncestor *TreeNode = nil

func postOrder(root, p, q *TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	}

	findP, findQ := postOrder(root.Left, p, q)
	if findP && findQ {
		if lastAncestor == nil { // 如果前方没有记录 则记录该点
			lastAncestor = root
		}
		return findP, findQ
	}

	resP, resQ := postOrder(root.Right, p, q)
	findP = findP || resP
	findQ = findQ || resQ
	if findP && findQ {
		if lastAncestor == nil {
			lastAncestor = root
		}

		return findP, findQ
	}

	if root.Val == p.Val {
		findP = true
	} else if root.Val == q.Val {
		findQ = true
	}

	if findP && findQ {
		if lastAncestor == nil {
			lastAncestor = root
		}

		return true, true
	}

	return findP, findQ
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	postOrder(root, p, q)

	res := lastAncestor
	lastAncestor = nil
	return res
}
