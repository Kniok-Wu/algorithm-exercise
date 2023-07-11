package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}

	str := strings.Builder{}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) != 0 {
			if stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	str.Write(stack)
	return str.String()
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
