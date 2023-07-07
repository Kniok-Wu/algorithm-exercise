package main

import (
	"fmt"
	"strings"
)

func FillIn(s *strings.Builder, str string, reverse bool) {
	if reverse {
		for i := len(str) - 1; i >= 0; i-- {
			s.WriteByte(str[i])
		}
	} else {
		for i := 0; i < len(str); i++ {
			s.WriteByte(str[i])
		}
	}
}

func reverseStr(s string, k int) string {
	str := &strings.Builder{}
	for i := 0; i < len(s); i += k {
		var input string
		if i+k <= len(s) {
			input = s[i : i+k]
		} else {
			input = s[i:]
		}
		var reverse bool
		if (i/k)%2 == 0 {
			reverse = true
		} else {
			reverse = false
		}
		FillIn(str, input, reverse)
	}

	return str.String()
}

//func reverseStr(s string, k int) string {
//	str := strings.Builder{}
//
//	i := 0
//	for ; i+k <= len(s); i += k {
//		if (i/k)%2 == 0 {
//			for j := k - 1; j >= 0; j-- { // 反转
//				str.WriteByte(s[i+j])
//			}
//
//		} else {
//			for j := 0; j < k; j++ { // 顺序写入
//				str.WriteByte(s[i+j])
//			}
//		}
//	}
//
//	if (i/k)%2 == 0 { // 等待反转
//		for j := len(s) - 1; j >= i; j-- {
//			str.WriteByte(s[j])
//		}
//
//	} else { // 已经反转过了
//		for i < len(s) {
//			str.WriteByte(s[i])
//			i += 1
//		}
//	}
//
//	return str.String()
//}

func main() {
	fmt.Println(reverseStr("a", 2))
}
