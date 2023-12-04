/**
 * @Time: 2023/11/26 18:59
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	cursor := 0

	for i := 0; i < len(chars); {
		j := i + 1
		for j < len(chars) && chars[j] == chars[i] {
			j++
		}
		count := j - i
		chars[cursor] = chars[i]
		cursor += 1

		if count > 1 {
			numStr := strconv.Itoa(count)
			for k := 0; k < len(numStr); k++ {
				chars[cursor] = numStr[k]
				cursor += 1
			}
		}
		i = j
	}

	return cursor
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c'}
	fmt.Println(compress(chars))
	fmt.Printf("%s\n", chars)
}
