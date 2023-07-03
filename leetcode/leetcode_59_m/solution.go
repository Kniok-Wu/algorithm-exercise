package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	// 初始化数组
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	result[0][0] = 1
	step := 1
	x, y := 0, 0
	for boundary := n - 1; boundary > -1; boundary-- {
		for i := 0; i <= boundary && y+step > -1 && y+step < n; i++ {
			result[x][y+step] = result[x][y] + 1
			y += step
		}

		for i := 0; i < boundary && x+step > 0 && x+step < n; i++ {
			result[x+step][y] = result[x][y] + 1
			x += step
		}

		step *= -1
	}

	return result
}

func main() {
	fmt.Println(generateMatrix(3))
}
