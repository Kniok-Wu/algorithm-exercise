package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) (bool, string) {
	for i := 1; i <= len(s)/2; i++ {
		subStr := s[:i]

		if len(s)%i != 0 {
			continue
		}

		j := i
		for ; j+i <= len(s); j += i {
			if s[j:j+i] != subStr {
				break
			}
		}
		if j == len(s) {
			return true, subStr
		}
	}
	return false, ""
}

func main() {
	//s := strings.Builder{}
	//for i := 0; i < 9999; i++ {
	//	s.WriteString("a")
	//}
	//s.WriteString("c")

	fmt.Println(repeatedSubstringPattern("abab"))
}
