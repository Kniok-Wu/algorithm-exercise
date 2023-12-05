package main

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	head := make([]int, 201)

	if len(nums) <= 1 {
		return result
	}

	var dfs func(int)

	records := make([]int, 0)

	dfs = func(start int) {
		if start == len(nums) {
			if len(records) >= 2 {
				result = append(result, append([]int{}, records...))
			}

			if len(records) > 0 {
				records = records[:len(records)-1]
			}

			return
		}

		i := start
		for i < len(nums) {
			// 如果records中为空 则直接添加一个数
			if len(records) == 0 {
				var headIdx int
				if nums[i] >= 0 {
					headIdx = nums[i]
				} else {
					headIdx = nums[i]
				}
				if head[headIdx] == 0 {
					head[headIdx] = 1
					records = append(records, nums[i])
				}
				i++
			} else {
				// 如果下一个数小最后一个数 直接跳过
				for i < len(nums) && nums[i] < records[len(records)-1] {
					i++
				}

				if i < len(nums) {
					records = append(records, nums[i])
					i++
				}
			}

			dfs(i)

			for ; i < len(nums) && nums[i-1] == nums[i]; i++ {
			}
		}
	}

	dfs(0)

	return result
}

func main() {
	result := findSubsequences([]int{1, 2, 1, 1, 1})
	for _, item := range result {
		fmt.Println(item)
	}
}
