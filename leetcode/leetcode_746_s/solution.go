/**
 * @Time: 2023/12/8 14:46
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	for i := 2; i < n; i++ {
		if cost[i-2] > cost[i-1] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}

	if cost[n-1] > cost[n-2] {
		return cost[n-2]
	}

	return cost[n-1]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
