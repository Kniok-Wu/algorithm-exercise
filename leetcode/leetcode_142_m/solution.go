package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	fast, slow := newHead, newHead

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = newHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
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

//func DisplayList(node *ListNode) {
//	for node != nil {
//		fmt.Printf("%d ", node.Val)
//		node = node.Next
//	}
//}

func main() {
	l := GenerateList([]int{1})
	//var tail *ListNode = l
	//for ; tail.Next != nil; tail = tail.Next {
	//}
	//
	//tail.Next = l.Next.Next.Next

	fmt.Println(detectCycle(l).Val)
}
