/**
 * @Time: 2023/12/8 16:03
 * @Author: kniokwu@gmail.com
 * @File: solution.md.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 0 {
		obstacleGrid[0][0] = 1
	} else {
		obstacleGrid[0][0] = 0
	}

	// 1. 对第一层进行遍历
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
			continue
		}

		obstacleGrid[0][j] += obstacleGrid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			if j-1 >= 0 {
				obstacleGrid[i][j] += obstacleGrid[i][j-1]
			}

			obstacleGrid[i][j] += obstacleGrid[i-1][j]
		}
	}

	return obstacleGrid[m-1][n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}))
}
