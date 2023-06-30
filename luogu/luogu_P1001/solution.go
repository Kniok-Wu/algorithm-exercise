package main

import (
	"fmt"
	"math"
)

func PrintMatrix(matrix [][]int) {
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			fmt.Printf("%2d ", matrix[x][y])
		}
		fmt.Println()
	}
}

func main() {
	var matrix [][]int

	var b_x, b_y, h_x, h_y int

	fmt.Scan(&b_x, &b_y, &h_x, &h_y)

	// 初始化棋盘 额外添加一层pad
	for i := 0; i < b_x+2; i++ {
		matrix = append(matrix, make([]int, b_y+2))
	}

	h_x += 1
	h_y += 1

	for x := 1; x < b_x+2; x++ {
		for y := 1; y < b_y+2; y++ {
			abs_x := math.Abs(float64(x - h_x)) // 由于添加了一圈pad 因此马坐标均 +1
			abs_y := math.Abs(float64(y - h_y))
			// 把是马点的位置设置为 -1
			if (abs_x+abs_y == 3 && math.Abs(abs_x-abs_y) != 3) || (x == h_x && y == h_y) {
				matrix[x][y] = -1
			}
		}
	}

	matrix[1][1] = 1
	//PrintMatrix(matrix)
	//fmt.Println()

	for x := 1; x < b_x+2; x++ {
		for y := 1; y < b_y+2; y++ {
			//PrintMatrix(matrix)
			//fmt.Println(fmt.Sprintf("当前点(%d, %d): %d\n", x, y, matrix[x][y]))
			// 如果这个点不能走 则设置为0
			if matrix[x][y] == -1 {
				matrix[x][y] = 0
				continue
			}

			matrix[x][y] += matrix[x][y-1] + matrix[x-1][y]
		}
	}

	fmt.Println(matrix[b_x+1][b_y+1])
}
