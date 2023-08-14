package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) <= 1 {
		result = append(result, []int{})
		if len(nums) == 1 {
			result = append(result, nums)
		}
		return result
	}

	sort.Ints(nums)

	var backTracking func([]int) [][]int

	backTracking = func(subNums []int) (subResult [][]int) {
		subResult = append(subResult, []int{})
		if len(subNums) <= 1 {
			if len(subNums) == 1 {
				subResult = append(subResult, append([]int{}, subNums...))
			}

			return
		}

		used := make(map[int]int)
		for i := 0; i < len(subNums); i++ {
			if used[subNums[i]] == 0 {
				for _, each := range backTracking(subNums[i+1:]) {
					subResult = append(subResult, append([]int{subNums[i]}, each...))
				}
				used[subNums[i]] = 1
			}
		}

		return
	}

	for _, item := range backTracking(nums) {
		result = append(result, append([]int{}, item...))
	}

	return result
}

func main() {
	for _, each := range subsetsWithDup([]int{4, 4, 4, 1, 4}) {
		fmt.Println(each)
	}
}
