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
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = nums[i]

		if last2 := i - 2; last2 >= 0 {
			dp[i] += dp[last2]
		}

		if last1 := i - 1; last1 >= 0 && dp[last1] > dp[i] {
			dp[i] = dp[last1]
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
