package main

import "fmt"

type Queue struct {
	q   []int
	cap int
}

func (q *Queue) Push(val int) {
	i := q.cap - 1
	for ; i >= 0; i-- {
		if q.q[i] >= val {
			break
		}
	}

	q.q = append(q.q[:i+1], val)
	q.cap = i + 2
}

func (q *Queue) Pop() {
	if len(q.q) == 1 {
		q.q = q.q[:0]
	} else {
		q.q = q.q[1:]
	}
	q.cap -= 1
}

func (q *Queue) GetMaxValue() int {
	return q.q[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	q := &Queue{
		q:   make([]int, 0),
		cap: 0,
	}

	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}

	result := []int{q.GetMaxValue()}

	for i := 1; i < len(nums) && i+k <= len(nums); i++ {
		if nums[i-1] == q.GetMaxValue() {
			q.Pop()
		}
		q.Push(nums[i+k-1])

		result = append(result, q.GetMaxValue())
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4))
}
