/**
 * @Time: 2023/12/15 13:30
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	if sum < abs(target) {
		return 0
	}

	dp := make([]int, sum*2+1)
	center := sum

	dp[center-nums[0]] += 1
	dp[center+nums[0]] += 1

	for i := 1; i < len(nums); i++ {
		newDp := make([]int, sum*2+1)
		for j := sum * 2; j >= 0; j-- {
			if dp[j] == 0 {
				continue
			} else {
				newDp[j+nums[i]] += dp[j]
				newDp[j-nums[i]] += dp[j]
			}
		}
		dp = newDp
	}

	return dp[center+target]
}

func findTargetSumWaysII(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	if (sum+target)%2 != 0 || sum < abs(target) {
		return 0
	}

	capacity := (sum + target) / 2

	dp := make([]int, capacity+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := capacity; j >= 0; j-- {
			if dp[j] == 0 || j+nums[i] > capacity {
				continue
			} else {
				dp[j+nums[i]] += dp[j]
			}
		}
	}

	return dp[capacity]
}

func main() {
	fmt.Println(findTargetSumWaysII([]int{1, 1, 1, 1, 1}, 3))
}
