/**
 * @Time: 2023/12/5 14:17
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	records := make([]int, 0, len(nums))
	result := make([][]int, 0)
	var dfs func(int)
	dfs = func(order int) {
		if len(records) == len(nums) {
			result = append(result, append([]int{}, records...))
			return
		}

		exists := make(map[int]struct{})

		for i := 0; i < len(nums); i++ {
			if _, ok := exists[nums[i]]; ok || nums[i] == 11 {
				continue
			}

			exists[nums[i]] = struct{}{}
			records = append(records, nums[i])
			nums[i] = 11
			dfs(order + 1)
			nums[i] = records[len(records)-1]
			records = records[:len(records)-1]
		}
	}

	dfs(0)

	return result
}

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(nums))
}
