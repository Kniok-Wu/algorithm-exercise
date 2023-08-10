package main

import "fmt"

func BackTracing(left []int, n, k int) [][]int {
	// 1. 不存在额外的数字了
	// 2. 期望的结果小于0
	// 3. 剩余可用数字为0
	result := make([][]int, 0)

	for i := 0; i < len(left); i++ {
		n_ := n - left[i] // 计算出还差的值

		if n_ == 0 { // 等于零 说明计算到期望值
			if k-1 == 0 { // 同时元素个数也满足 则假如结果中
				result = append(result, []int{left[i]})
			}
			break // 由于元素唯一 所以直接跳出
		} else if n_ > left[i] && k-1 > 0 { // 所差值大于零 且所差值大于当前值(因为后序值越来越大) 且仍包含所需元素
			childResult := BackTracing(left[i+1:], n-left[i], k-1)
			for _, each := range childResult {
				result = append(result, append([]int{left[i]}, each...))
			}
		}
	}

	return result
}

func combinationSum3(k int, n int) [][]int {
	/**
	k: k个数
	n: 和为n
	*/
	left := make([]int, 9)
	for i := 0; i < 9; i++ {
		left[i] = i + 1
	}

	return BackTracing(left, n, k)
}

func main() {
	for _, each := range combinationSum3(4, 1) {
		fmt.Println(each)
	}
}
