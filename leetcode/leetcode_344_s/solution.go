package main

import (
	"fmt"
	"strings"
)

func reverseString(s []byte) {
	last := len(s) - 1
	start := 0
	for start < last {
		s[start], s[last] = s[last], s[start]
		start += 1
		last -= 1
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	r := strings.Builder{}
	r.Write(s)
	fmt.Println(r.String())
}
