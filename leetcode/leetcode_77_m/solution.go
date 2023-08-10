package main

import "fmt"

var result [][]int

func BackTracking(base []int, left []int, k int) {
	if len(base)+len(left) < k {
		return
	}

	if len(base) == k {
		result = append(result, append([]int{}, base...))
	}

	for i := 0; i < len(left); i++ {
		BackTracking(append(base, left[i]), left[i+1:], k)
	}

}

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = i + 1
	}
	BackTracking([]int{}, left, k)

	return result
}

func main() {
	n, k := 5, 4

	for _, each := range combine(n, k) {
		fmt.Println(each)
	}
}
