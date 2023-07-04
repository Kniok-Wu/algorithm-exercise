package main

import "fmt"

type MyLinkedList struct {
	Val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  1001, // <= 1000
		next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this // skip head node
	for i := 0; i <= index && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{
		Val:  val,
		next: this.next,
	}
	this.next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &MyLinkedList{
		Val:  val,
		next: nil,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this
	i := 0
	for ; i < index && cur != nil; i++ {
		cur = cur.next
	} // 查询直到末尾 或者 找到目标元素前一个

	if i-1 == index && cur == nil { // 查询到最后一个元素
		this.AddAtTail(val)
	} else if cur != nil { // 查询到目标元素
		temp := &MyLinkedList{
			Val:  val,
			next: cur.next,
		}
		cur.next = temp
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	i := 0
	for ; i < index && cur != nil; i++ { // 查询到前一个 或者 最后一个元素
		cur = cur.next
	}

	if cur != nil && cur.next != nil { // 该元素和该元素的后一个元素不为空
		cur.next = cur.next.next
	}
}

func main() {
	head := Constructor()
	head.AddAtHead(1)
	head.AddAtTail(3)
	head.AddAtIndex(1, 2)
	fmt.Println(head.Get(1))
	head.DeleteAtIndex(1)
	fmt.Println(head.Get(1))
	fmt.Println()
}
