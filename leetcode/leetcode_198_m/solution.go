/**
 * @Time: 2024/1/8 14:53
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func rob(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]

		last1 := i - 1
		last2 := i - 2

		if last2 >= 0 {
			dp[i] += dp[last2]
		}

		if last1 >= 0 && dp[i] < dp[last1] {
			dp[i] = dp[last1]
		}
	}

	return dp[len(nums)-1]
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
