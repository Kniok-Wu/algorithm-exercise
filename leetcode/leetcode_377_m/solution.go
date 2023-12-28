/**
 * @Time: 2023/12/28 21:35
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1

	for total := 0; total <= target; total++ {
		for _, num := range nums {
			if total >= num {
				dp[total] += dp[total-num]
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
