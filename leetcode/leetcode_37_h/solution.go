/**
 * @Time: 2023/12/7 21:13
 * @Author: kniokwu@gmail.com
 * @File: solution.md.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func boxNum(x, y int) int {
	var line int
	if y < 3 {
		line = 1
	} else if y > 5 {
		line = 3
	} else {
		line = 2
	}

	box := (line-1)*3 - 1
	if x < 3 {
		return box + 1
	} else if x > 5 {
		return box + 3
	} else {
		return box + 2
	}
}

func solveSudoku(board [][]byte) {
	lines := make([][]bool, 9)
	cols := make([][]bool, 9)
	boxes := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		lines[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
		boxes[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				key := board[i][j] - '1'
				lines[i][key] = true
				cols[j][key] = true
				boxes[boxNum(i, j)][key] = true
			}
		}
	}

	finished := false

	var dfs func(int, int)
	dfs = func(line, col int) {
		// 1. 该处是否已经有数字 如果有 则直接进入下一次迭代
		if board[line][col] != '.' {
			// 1. 向下移动
			if col+1 > 8 && line < 8 {
				// 2.1 向下移动
				dfs(line+1, 0)
			}

			// 2. 向右移动
			if col+1 <= 8 {
				dfs(line, col+1)
			}

			// 3. 已经完成布局
			if col == 8 && line == 8 {
				finished = true
				return
			}
		}

		// 2. 需要自行放置
		if board[line][col] == '.' {
			for num := byte('1'); num <= byte('9'); num++ {
				// 2.1 可以放置 则尝试进行下一次迭代
				box := boxNum(line, col)
				key := num - '1'
				if !lines[line][key] && !cols[col][key] && !boxes[box][key] {
					board[line][col] = num
					lines[line][key] = true
					cols[col][key] = true
					boxes[box][key] = true

					// 1. 向下移动
					if col+1 > 8 && line < 8 {
						// 2.1 向下移动
						dfs(line+1, 0)
					}

					// 2. 向右移动
					if col+1 <= 8 {
						dfs(line, col+1)
					}

					// 3. 已经完成布局
					if col == 8 && line == 8 || finished {
						finished = true
						return
					}

					board[line][col] = '.'
					lines[line][key] = false
					cols[col][key] = false
					boxes[box][key] = false
				}
			}
		}
	}

	dfs(0, 0)
}

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}
		fmt.Println()
	}
}
