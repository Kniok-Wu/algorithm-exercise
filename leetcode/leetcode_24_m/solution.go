package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre := newHead
	first := newHead.Next

	for first != nil && first.Next != nil {
		pre.Next = first.Next
		first.Next = pre.Next.Next
		pre.Next.Next = first
		pre = first
		first = pre.Next
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
	DisplayList(swapPairs(GenerateList([]int{1, 2, 3})))
}
