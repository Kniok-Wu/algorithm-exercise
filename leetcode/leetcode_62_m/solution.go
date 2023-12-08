/**
 * @Time: 2023/12/8 15:48
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func uniquePaths(m int, n int) int {
	pathes := make([][]int, m)
	for i := 0; i < m; i++ {
		pathes[i] = make([]int, n)
	}

	pathes[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				pathes[i][j] += pathes[i-1][j]
			}

			if j-1 >= 0 {
				pathes[i][j] += pathes[i][j-1]
			}
		}
	}

	return pathes[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
