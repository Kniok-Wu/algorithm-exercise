package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	result[0][0] = 1 // 初始化数组
	step := 1
	x, y := 0, 0
	for boundary := n - 1; boundary > -1; boundary-- {
		/*
				以 3*3 矩阵为例，每次填入2条边
				0 1 2
				7 8 3
				6 5 4
				第一次 第一条边 3个元素 步进 1
			          第二条边 2个元素 步进 1
				第二次 第一条边 2个元素 步进 -1
				      第二条边 1个元素 步进 -1
			    第二次 第一条边 1个元素 步进 -1
					  第二条边 0个元素 步进 -1
				可以很清楚的看见，boundary初始为n-1，每轮递减，每轮中第一条边遍历边界为x<=boundry，每轮中第二条边遍历边界为x<boundry
				每完成一轮 step转为-1(向上/向左)
				同时为什么 i <= boundary && y+step > -1 && y+step < n，是因为第一组里 result[0][0]没法向前取值，因此添加了额外的限制条件保证一致性。
		*/
		for i := 0; i <= boundary && y+step > -1 && y+step < n; i++ {
			result[x][y+step] = result[x][y] + 1
			y += step
		}

		for i := 0; i < boundary; i++ {
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
