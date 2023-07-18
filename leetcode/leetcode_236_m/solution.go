package leetcode_236_m

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ancestorP = make([]*TreeNode, 0)
var ancestorQ = make([]*TreeNode, 0)

func Check(subRootP, subRootQ, p, q int) {
	if ancestorP[len(ancestorP)-1].Val != p { // 子树中没有该节点 则恢复到左子树的根节点 准备右子树
		ancestorP = ancestorP[:subRootP]
	}

	if ancestorQ[len(ancestorQ)-1].Val != q { // 子树中没有该节点 则恢复到左子树的根节点 准备右子树
		ancestorQ = ancestorQ[:subRootQ]
	}
}

func Traversal(root *TreeNode, p, q int) {
	if root == nil {
		return
	}

	count := 0

	lp := len(ancestorP)
	if !(lp > 0 && ancestorP[lp-1].Val == p) { // 没有记录到目标元素
		ancestorP = append(ancestorP, root)
		lp += 1
		count += 1
	}
	lq := len(ancestorQ)
	if !(lq > 0 && ancestorQ[lq-1].Val == q) { // 没有记录到目标元素
		ancestorQ = append(ancestorQ, root)
		lq += 1
		count += 1
	}

	if count == 0 { // 两个元素均找到了
		return
	}

	Traversal(root.Left, p, q)
	Check(lp, lq, p, q)
	lp = len(ancestorP)
	lq = len(ancestorQ)
	Traversal(root.Right, p, q)
	Check(lp, lq, p, q)
}

func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	Traversal(root, p.Val, q.Val)

	var res *TreeNode = nil
	for i := len(ancestorP) - 1; i >= 0; i -= 1 {
		for j := len(ancestorQ) - 1; j >= 0; j -= 1 {
			if ancestorQ[j].Val == ancestorP[i].Val {
				res = ancestorP[i]
				goto RETURN
			}
		}
	}

RETURN:
	ancestorP = make([]*TreeNode, 0)
	ancestorQ = make([]*TreeNode, 0)

	return res
}
