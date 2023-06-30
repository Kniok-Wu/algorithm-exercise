package main

import (
	"fmt"
)

func InsertBracket(record *map[string]bool, base string, num int) {
	if num == 0 {
		return
	}

	for i := 0; i < len(base); i++ {
		str := base[:i] + "()" + base[i:]
		if _, ok := (*record)[str]; !ok {
			(*record)[str] = true
			InsertBracket(record, str, num-1)
		}
	}

	return
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	record := map[string]bool{
		"()": true,
	}
	InsertBracket(&record, "()", n-1)
	result := make([]string, 0)
	for k, _ := range record {
		if len(k) == 2*n {
			result = append(result, k)
		}
	}
	return result
}

func generateParenthesis_2(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}

func main() {
	fmt.Println(generateParenthesis_2(20))
}
