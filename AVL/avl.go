package AVLTree

type AVLNode struct {
	val                    int
	depth                  int
	lchild, rchild, parent *AVLNode
}

type AVLTree struct {
	root *AVLNode
}

func Constructor() *AVLTree {
	return &AVLTree{}
}

// Search 搜索树中是否存在该元素
func (t *AVLTree) Search(val int) bool {
	for cur := t.root; cur != nil; {
		if cur.val == val {
			return true
		} else if cur.val > val {
			cur = cur.lchild
		} else {
			cur = cur.rchild
		}
	}
	return false
}

// getBalanceFactor 获取一个节点的平衡因子
func (n *AVLNode) getBalanceFactor() int {
	leftHeight := 0
	rightHeight := 0
	if n.lchild != nil {
		leftHeight = n.lchild.depth
	}
	if n.rchild != nil {
		rightHeight = n.rchild.depth
	}

	return rightHeight - leftHeight
}

// Insert 添加一个元素进入平衡二叉搜索树中
func (t *AVLTree) Insert(val int) {
	node := &AVLNode{
		val:    val,
		depth:  1,
		lchild: nil,
		rchild: nil,
		parent: nil,
	}

	if t.root == nil { // 空树 则添加至根节点
		t.root = node
		return
	}

	if !t.Search(val) { // 不存在 则添加该节点
		cur := t.root
		for cur != nil {
			node.parent = cur
			if val < cur.val {
				if cur.lchild == nil {
					cur.lchild = node
					updateHeight(t.root)
					break
				}
				cur = cur.lchild
			} else {
				if cur.rchild == nil {
					cur.rchild = node
					updateHeight(t.root)
					break
				}
				cur = cur.rchild
			}
		}

		// 检查平衡因子
		adjust := false
		for cur = node; cur != nil; cur = cur.parent {
			if cur.getBalanceFactor() == -2 {
				if cur.lchild != nil && cur.lchild.getBalanceFactor() == -1 {
					t.LLAdjust(cur)
					adjust = true
				}
			}

			if cur.getBalanceFactor() == 2 {
				if cur.rchild != nil && cur.rchild.getBalanceFactor() == 1 {
					t.RRAdjust(cur)
					adjust = true
				}
			}

			if cur.getBalanceFactor() == 2 {
				if cur.rchild != nil && cur.rchild.getBalanceFactor() == -1 {
					t.RLAdjust(cur)
					adjust = true
				}
			}

			if cur.getBalanceFactor() == -2 {
				if cur.rchild != nil && cur.rchild.getBalanceFactor() == 1 {
					t.LRAdjust(cur)
					adjust = true
				}
			}

			if adjust {
				updateHeight(t.root)
				break
			}
		}
	}
}

func (t *AVLTree) Remove(val int) {
	cur := t.root
	for cur != nil {
		if cur.val > val {
			cur = cur.lchild
		} else if cur.val < val {
			cur = cur.rchild
		} else if cur.val == val {
			break
		}
	}

	if cur != nil {
		if cur.parent == nil { // 删除整根树
			t.root = nil
			return
		}

		cur.parent.depth = 1
		if cur.parent.val > val {
			cur.parent.lchild = nil
			if cur.parent.rchild != nil {
				cur.parent.depth += cur.parent.rchild.depth
			}
		} else {
			cur.parent.rchild = nil
			if cur.parent.lchild != nil {
				cur.parent.depth += cur.parent.lchild.depth
			}
		}

		for cur = cur.parent; cur != nil; cur = cur.parent {
			if cur.getBalanceFactor() == -2 {
				if cur.lchild != nil && cur.lchild.getBalanceFactor() == -1 {
					t.LLAdjust(cur)
					cur = cur.parent
				}
			}
			if cur.getBalanceFactor() == -2 {
				if cur.lchild != nil && cur.lchild.getBalanceFactor() == 1 {
					t.LRAdjust(cur)
					cur = cur.parent
				}
			}
			if cur.getBalanceFactor() == 2 {
				if cur.rchild != nil && cur.rchild.getBalanceFactor() == 1 {
					t.RRAdjust(cur)
					cur = cur.parent
				}
			}

			if cur.getBalanceFactor() == 2 {
				if cur.rchild != nil && cur.rchild.getBalanceFactor() == 1 {
					t.RLAdjust(cur)
					cur = cur.parent
				}
			}
		}
	}
}

func (t *AVLTree) LLAdjust(unbalanceHead *AVLNode) {
	a := unbalanceHead
	b := a.lchild
	if b.rchild != nil {
		b.rchild.parent = a
	}
	b.rchild, a.lchild = a, b.rchild
	b.parent = a.parent
	if b.parent != nil {
		b.parent.lchild = b
	} else {
		t.root = b
	}
	a.parent = b
}

func (t *AVLTree) RRAdjust(unbalanceHead *AVLNode) {
	a := unbalanceHead
	b := unbalanceHead.rchild
	if b.lchild != nil {
		b.lchild.parent = a
	}
	a.rchild, b.lchild = b.lchild, a
	b.parent = a.parent
	if a.parent != nil {
		a.parent.rchild = b
	} else {
		t.root = b
	}
	a.parent = b
}

func (t *AVLTree) LRAdjust(unbalanceHead *AVLNode) {
	a := unbalanceHead
	b := unbalanceHead.lchild
	b.rchild.parent = a
	b.parent = b.rchild
	a.lchild, b.rchild.lchild = b.rchild, b
	b.rchild = nil

	t.LLAdjust(unbalanceHead)
}

func (t *AVLTree) RLAdjust(unbalanceHead *AVLNode) {
	a := unbalanceHead
	b := unbalanceHead.lchild
	b.lchild.parent = a
	b.parent = b.lchild
	a.rchild, b.lchild.rchild = b.lchild, b
	b.lchild = nil

	t.RRAdjust(unbalanceHead)
}

func updateHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}

	var height int
	lHeight := updateHeight(node.lchild)
	rHeight := updateHeight(node.rchild)
	if lHeight >= height {
		height = lHeight + 1
	}
	if rHeight >= height {
		height = rHeight + 1
	}
	node.depth = height
	return height
}
