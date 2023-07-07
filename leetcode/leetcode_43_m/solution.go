package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	length := (len(num1)+1)*(len(num2)+1) - 2
	result := make([]int, length)

	offset1, offset2 := -1, -1
	var flag int = 0 // 进位
	for i := len(num1) - 1; i >= 0; i-- {
		offset1 += 1
		b1 := int(num1[i]) - int('0')
		for j := len(num2) - 1; j >= 0; j-- {
			offset2 += 1
			b2 := int(num2[j]) - int('0')
			temp := b1*b2 + result[offset1+offset2] + flag
			if temp >= 10 {
				result[offset1+offset2] = temp % 10
				flag = temp / 10
			} else {
				result[offset1+offset2] = temp
				flag = 0
			}
		}
		if flag != 0 {
			result[offset1+offset2+1] += flag
			flag = 0
		}
		offset2 = -1
	}

	end := length - 1

	for end > 0 && result[end] == 0 {
		end -= 1
	}

	s := strings.Builder{}

	for ; end >= 0; end-- {
		s.WriteByte(byte(result[end] + int('0')))
	}

	return s.String()
}

func main() {
	fmt.Println(multiply("9", "99"))
}
