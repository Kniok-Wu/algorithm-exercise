package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums, []int{}}
	}

	var backTracking func([]int) [][]int

	backTracking = func(subNums []int) [][]int {
		mid := len(subNums) / 2
		stack := make([][]int, 0)

		left := subNums[:mid]
		right := subNums[mid:]

		if len(left) == 1 {
			stack = append(stack, append([]int{}, left...))
		} else if len(left) > 1 {
			for _, item := range backTracking(left) {
				stack = append(stack, append([]int{}, item...))
			}
		}

		n := len(stack)
		if len(right) == 1 {
			stack = append(stack, append([]int{}, right...))
			for i := 0; i < n; i++ {
				stack = append(stack, append(append([]int{}, stack[i]...), right...))
			}
		} else if len(right) > 1 {
			for _, item := range backTracking(right) {
				stack = append(stack, append([]int{}, item...))
				for i := 0; i < n; i++ {
					stack = append(stack, append(append([]int{}, item...), stack[i]...))
				}
			}
		}

		return stack
	}

	result := append(backTracking(nums), []int{})

	return result
}

func main() {
	for _, item := range subsets([]int{1, 2, 3, 4}) {
		fmt.Println(item)
	}
}
