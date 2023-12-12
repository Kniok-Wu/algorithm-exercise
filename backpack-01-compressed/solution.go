/**
 * @Time: 2023/12/12 10:51
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

//func main() {
//	var n, capacity int
//	fmt.Scan(&n, &capacity)
//
//	values := make([]int, n)
//	weights := make([]int, n)
//
//	for i := 0; i < n; i++ {
//		fmt.Scan(&weights[i], &values[i])
//	}
//
//	dp := make([]int, capacity+1)
//
//	for i := weights[0]; i <= capacity; i++ {
//		dp[i] = values[0]
//	}
//
//	for i := 1; i < n; i++ {
//		newDp := make([]int, capacity+1)
//		for j := 0; j <= capacity; j++ {
//			// 1. 如果无法选择当前物件 直接继承上一物件
//			if j < weights[i] {
//				newDp[j] = dp[j]
//			}
//
//			// 2. 可以选择当前物件
//			if j >= weights[i] {
//				curValue := values[i] + dp[j-weights[i]]
//				if curValue > dp[j] {
//					newDp[j] = curValue
//				} else {
//					newDp[j] = dp[j]
//				}
//			}
//		}
//		dp = newDp
//	}
//
//	fmt.Println(dp[capacity])
//}

func main() {
	var n, capacity int
	fmt.Scan(&n, &capacity)

	values := make([]int, n)
	weights := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i], &values[i])
	}

	dp := make([]int, capacity+1)

	for i := weights[0]; i <= capacity; i++ {
		dp[i] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := capacity; j >= 0; j-- {
			if j < weights[i] {
				break
			}

			if j >= weights[i] {
				curValue := values[i] + dp[j-weights[i]]
				if curValue > dp[j] {
					dp[j] = curValue
				}
			}
		}
	}

	fmt.Println(dp[capacity])
}
