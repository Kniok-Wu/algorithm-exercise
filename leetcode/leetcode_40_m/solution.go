package main

import "fmt"

func FastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[0]

	i, j := 0, len(nums)-1

	for i != j {
		for nums[j] >= pivot && i != j {
			j--
		}

		nums[i] = nums[j]

		for nums[i] <= pivot && i != j {
			i++
		}

		nums[j] = nums[i]
	}

	nums[i] = pivot

	FastSort(nums[:i])
	FastSort(nums[i+1:])
}

func BackTracking(target int, left []int) [][]int {
	result := make([][]int, 0)

	if !(target <= 0 || len(left) == 0) {
		for i := 0; i < len(left); {
			if target_ := target - left[i]; target_ > 0 {
				if i+1 != len(left) {
					for _, each := range BackTracking(target_, left[i+1:]) {
						result = append(result, append([]int{left[i]}, each...))
					}
				}
			} else if target_ == 0 { // 找到结果
				result = append(result, []int{left[i]})
			}

			for i = i + 1; i < len(left) && left[i] == left[i-1]; i++ {
			}
		}
	}

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	FastSort(candidates) // 对数组进行排序
	return BackTracking(target, candidates)
}

func main() {
	for _, each := range combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8) {
		fmt.Println(each)
	}
}
