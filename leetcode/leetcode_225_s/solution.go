package main

import "fmt"

type MyStack struct {
	queue_1 []int
	queue_2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue_1: make([]int, 0),
		queue_2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	if len(this.queue_1) == 0 {
		this.queue_2 = append(this.queue_2, x)
	} else {
		this.queue_1 = append(this.queue_1, x)
	}
}

func (this *MyStack) Pop() int {
	res := -1
	if len(this.queue_1) == 0 {
		for len(this.queue_2) > 1 {
			this.queue_1 = append(this.queue_1, this.queue_2[0])
			this.queue_2 = this.queue_2[1:]
		}
		res = this.queue_2[0]
		this.queue_2 = make([]int, 0)
	} else {
		for len(this.queue_1) > 1 {
			this.queue_2 = append(this.queue_2, this.queue_1[0])
			this.queue_1 = this.queue_1[1:]
		}
		res = this.queue_1[0]
		this.queue_1 = make([]int, 0)
	}
	return res
}

func (this *MyStack) Top() int {
	res := -1
	if len(this.queue_1) == 0 {
		for len(this.queue_2) > 1 {
			this.queue_1 = append(this.queue_1, this.queue_2[0])
			this.queue_2 = this.queue_2[1:]
		}
		res = this.queue_2[0]
		this.queue_1 = append(this.queue_1, res)
		this.queue_2 = make([]int, 0)
	} else {
		for len(this.queue_1) > 1 {
			this.queue_2 = append(this.queue_2, this.queue_1[0])
			this.queue_1 = this.queue_1[1:]
		}
		res = this.queue_1[0]
		this.queue_2 = append(this.queue_2, res)
		this.queue_1 = make([]int, 0)
	}
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.queue_1) == 0 && len(this.queue_2) == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Top())
	fmt.Println(q.Top())
	fmt.Println(q.Empty())
}
