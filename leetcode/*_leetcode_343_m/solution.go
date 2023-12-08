/**
 * @Time: 2023/12/8 16:17
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0 // num = 0 的时候
	dp[1] = 1 // num = 1 的时候
	dp[2] = 1 // num = 1 的时候

	for i := 3; i < n+1; i++ {
		max := 0

		for j := 1; j <= i/2; j++ {
			num1, num2 := j, i-j
			if num1 < dp[num1] {
				num1 = dp[num1]
			}

			if num2 < dp[num2] {
				num2 = dp[num2]
			}

			if num1*num2 > max {
				max = num1 * num2
			}
		}

		dp[i] = max
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(3))
}
