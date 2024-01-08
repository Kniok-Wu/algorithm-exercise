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

	max := nums[0]

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		for j := i - 2; j >= 0; j-- {
			if dp[i] < nums[i-1]+dp[j] {
				dp[i] = nums[i-1] + dp[j]
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func main() {
	fmt.Println(rob([]int{1}))
}
