package main

import "fmt"

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	i, j, k := 0, len(nums)-1, len(nums)-1

	for i <= j {
		t_j := nums[j] * nums[j]
		t_i := nums[i] * nums[i]

		if t_i > t_j {
			result[k] = t_i
			i++
		} else {
			result[k] = t_j
			j--
		}
		k--
	}
	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-1, 0, 1, 2, 3}))
}
