/**
 * @Time: 2023/12/3 21:18
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func reverse(x int) int {
	minus := false
	if x < 0 {
		minus = true
		x *= -1
	}
	res := 0

	for x != 0 {
		temp := x % 10
		res = res*10 + temp
		if res < 0 || res > (1<<31) {
			return 0
		}

		x = (x - temp) / 10
	}

	if minus {
		res = res * -1
	}

	return res
}

func main() {
	fmt.Println(reverse(123))
}
