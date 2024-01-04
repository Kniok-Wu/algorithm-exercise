/**
 * @Time: 2024/1/4 10:47
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		for j := 0; j < len(wordDict) && !dp[i]; j++ {
			word := wordDict[j]
			if len(word) <= i {
				dp[i] = dp[i-len(word)] && s[:i][i-len(word):] == word
			}
		}
	}

	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
