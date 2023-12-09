/**
 * @Time: 2023/12/9 15:02
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
