/**
 * @Time: 2023/12/4 21:59
 * @Author: kniokwu@gmail.com
 * @File: solution.md.md.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	records := make([]int, 0, len(nums))

	var dfs func()
	dfs = func() {
		if len(records) == len(nums) {
			result = append(result, append([]int{}, records...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if nums[i] == 11 {
				continue
			}

			records = append(records, nums[i])
			nums[i] = 11
			dfs()
			nums[i] = records[len(records)-1]
			records = records[:len(records)-1]
		}
	}

	dfs()

	return result
}

func main() {
	fmt.Println(permute([]int{1}))
}
