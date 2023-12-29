/**
 * @Time: 2023/12/29 19:48
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import (
	"fmt"
)

func numSquares(n int) int {

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i

		for j := 1; j*j <= i; j++ {
			cur := dp[i-j*j] + 1
			if cur < dp[i] {
				dp[i] = cur
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(9999))
}
