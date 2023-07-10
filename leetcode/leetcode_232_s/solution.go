package main

import "fmt"

type MyQueue struct {
	stack_1 []int
	stack_2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack_1: make([]int, 0),
		stack_2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack_1 = append(this.stack_1, x)
}

func (this *MyQueue) Pop() int {
	res := this.Peek()
	this.stack_2 = this.stack_2[1:]
	return res
}

func (this *MyQueue) Peek() int {
	length := len(this.stack_1)
	for i := 0; i < length; i++ {
		this.stack_2 = append(this.stack_2, this.stack_1[0])
		this.stack_1 = this.stack_1[1:]
	}
	return this.stack_2[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack_1) == 0 && len(this.stack_2) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	//obj.Push(1)
	//obj.Push(2)
	//fmt.Println(obj.Peek())
	//fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
