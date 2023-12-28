/**
 * @Time: 2023/12/28 16:06
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func display(dp [][]int) {
	fmt.Print("  ")
	for i := 0; i < len(dp[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("  ")
	for i := 0; i < len(dp); i++ {
		fmt.Print(i)
		fmt.Println(dp[i])
	}
}

func findMaxForm(strs []string, m int, n int) int {
	countZero := make([]int, len(strs))
	countOne := make([]int, len(strs))

	for idx, str := range strs {
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				countZero[idx] += 1
			} else {
				countOne[idx] += 1
			}
		}
	}

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for item := 0; item < len(strs); item++ {
		for i := m; i >= 0; i-- {
			if i < countZero[item] {
				break
			}
			for j := n; j >= 0; j-- {
				if j < countOne[item] {
					break
				}

				newVal := dp[i-countZero[item]][j-countOne[item]] + 1
				if newVal > dp[i][j] {
					dp[i][j] = newVal
				}
			}
		}
		//display(dp)
		//fmt.Println()
	}

	return dp[m][n]
}

func main() {
	strs := []string{"11111", "100", "1101", "1101", "11000"}
	fmt.Println(findMaxForm(strs, 5, 7))
}
