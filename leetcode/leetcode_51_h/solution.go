/**
 * @Time: 2023/12/5 14:41
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import (
	"fmt"

	"strings"
)

func generateString(index int, n int) string {
	sb := strings.Builder{}
	for i := 0; i < index; i++ {
		sb.WriteByte('.')
	}
	sb.WriteByte('Q')
	for i := index + 1; i < n; i++ {
		sb.WriteByte('.')
	}
	return sb.String()
}

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	records := make([]string, 0, n)
	exists := make([]int, n*n)

	var dfs func(layer int)

	dfs = func(layer int) {
		if layer == n {
			result = append(result, append([]string{}, records...))
			return
		}

		for i := 0; i < n; i++ {
			idx := layer*n + i
			if exists[idx] != 0 {
				continue
			}

			lastExists := make([]int, 0)
			lastExists = append(lastExists, exists...)

			for j := layer; j < n; j++ {
				idx = j*n + i
				exists[idx] = -1
			}
			j, k := layer, i
			for l := 1; l < n; l++ {
				if j+l < n && k+l < n {
					idx = (j+l)*n + k + l
					exists[idx] = -1
				}

				if j+l < n && k-l >= 0 {
					idx = (j+l)*n + k - l
					exists[idx] = -1
				}

				if j-l >= 0 && k+l < n {
					idx = (j-l)*n + k + l
					exists[idx] = -1
				}

				if j-l >= 0 && k-l >= 0 {
					idx = (j-l)*n + k - l
					exists[idx] = -1
				}
			}
			records = append(records, generateString(i, n))
			dfs(layer + 1)
			records = records[:len(records)-1]
			for j = 0; j < n*n; j++ {
				exists[j] = lastExists[j]
			}
		}
	}

	dfs(0)

	return result
}

func main() {
	for _, each := range solveNQueens(20) {
		fmt.Println(each)
	}
}
