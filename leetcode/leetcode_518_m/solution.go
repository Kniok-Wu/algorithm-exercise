/**
 * @Time: 2023/12/28 19:28
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

// 1. 先遍历物品 再遍历背包
func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for total := coin; total <= amount; total++ {
			dp[total] += dp[total-coin]
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
