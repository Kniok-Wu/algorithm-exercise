/**
 * @Time: 2024/1/8 16:19
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	max := nums[0]

	if n == 1 {
		return max
	}

	if cur := subRob(nums[:n-1]); cur > max {
		max = cur
	}

	if cur := subRob(nums[1:]); cur > max {
		max = cur
	}

	return max
}

func subRob(nums []int) int {
	n := len(nums)
	max := nums[0]
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = nums[i]

		for j := i - 2; j >= 0; j-- {
			if nums[i]+dp[j] > dp[i] {
				dp[i] = nums[i] + dp[j]
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
