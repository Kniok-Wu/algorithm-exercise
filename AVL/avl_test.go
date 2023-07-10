package AVLTree

import (
	"fmt"
	"testing"
)

func firstTraverse(node *AVLNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.val)
	firstTraverse(node.lchild)
	firstTraverse(node.rchild)
}

func TestConstructor(t *testing.T) {
	head := Constructor()
	//nums := []int{7, 5, 8, 4, 6, 3}
	nums := []int{10, 5, 11, 3, 8, 4, 7, 6}
	for _, val := range nums {
		head.Insert(val)
	}
	firstTraverse(head.root)
	fmt.Println()
	head.Remove(10)
	firstTraverse(head.root)
	fmt.Println()
}
