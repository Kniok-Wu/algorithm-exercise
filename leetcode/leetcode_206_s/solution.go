package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{
		Val:  -1,
		Next: nil,
	}

	cur := head

	for cur != nil {
		next := newHead.Next
		newHead.Next = cur
		cur = cur.Next
		newHead.Next.Next = next
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
	DisplayList(reverseList(GenerateList([]int{1, 2, 3, 4, 5})))
}
