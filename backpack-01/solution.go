/**
 * @Time: 2023/12/11 14:52
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func main() {
	var n, capacity int
	fmt.Scan(&n, &capacity)

	values := make([]int, n)
	weights := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i], &values[i])
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, capacity+1)
	}

	for i := 0; i <= capacity; i++ {
		if i < weights[0] {
			dp[0][i] = []int{0, 0}
		} else {
			dp[0][i] = []int{weights[0], values[0]}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= capacity; j++ {
			// 1. 无法选择当前物件
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			}

			// 2. 可以选择当前物件
			if j >= weights[i] {
				dp[i][j] = []int{weights[i], values[i]}
				last := dp[i-1][j-weights[i]]

				for k := 0; k < 2; k++ {
					dp[i][j][k] += last[k]
				}

				if dp[i][j][1] < dp[i-1][j][1] {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	fmt.Println(dp[n-1][capacity][1])
}
