package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		-1,
		head,
	}

	last := newHead
	cur := newHead.Next

	for cur != nil {
		if cur.Val == val {
			last.Next = cur.Next
		} else {
			last = last.Next
		}
		cur = cur.Next
	}

	return newHead.Next
}

func GenerateList(nums []int) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}

	p := head

	for _, val := range nums {
		p.Next = &ListNode{
			val,
			nil,
		}
		p = p.Next
	}
	return head.Next
}

func DisplayList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
}

func main() {
	DisplayList(removeElements(GenerateList([]int{1, 2, 6, 3, 4, 5, 6}), 6))
}
