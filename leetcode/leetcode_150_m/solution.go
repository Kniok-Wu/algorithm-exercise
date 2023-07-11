package main

import (
	"fmt"
	"strconv"
)

func calc(pre, after int, sign string) int {
	switch sign {
	case "+":
		return pre + after
	case "-":
		return pre - after
	case "*":
		return pre * after
	default:
		return pre / after
	}
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	length := 0

	for i := 0; i < len(tokens); i++ {
		if val, err := strconv.Atoi(tokens[i]); err == nil {
			stack = append(stack, val)
			length += 1
		} else {
			res := calc(stack[length-2], stack[length-1], tokens[i])
			length -= 1
			stack = stack[:length]
			stack[length-1] = res
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
