/**
 * @Time: 2023/12/15 12:46
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, val := range stones {
		sum += val
	}

	capacity := int(sum / 2)
	dp := make([]int, capacity+1)

	for i := stones[0]; i <= capacity; i++ {
		dp[i] = stones[0]
	}

	for i := 1; i < len(stones); i++ {
		for j := capacity; j >= 0; j-- {
			if j < stones[i] {
				break
			}

			if j >= stones[i] {
				curVal := stones[i] + dp[j-stones[i]]
				if curVal > dp[j] {
					dp[j] = curVal
				}
			}
		}
	}

	fmt.Println(dp)

	return sum - dp[capacity]*2
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
