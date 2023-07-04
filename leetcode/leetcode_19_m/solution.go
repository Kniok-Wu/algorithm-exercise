package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	slow := newHead
	fast := newHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
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
	DisplayList(removeNthFromEnd(GenerateList([]int{1}), 1))
}
