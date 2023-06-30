package main

import (
	"fmt"
	"strconv"
)

var letterMap = map[int][]string{
	0: []string{""},
	1: []string{""},
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {

	length := len(digits)

	if length == 0 {
		return nil
	}

	return Search(digits)
}

func Combine(arr_1, arr_2 []string) []string {
	res := make([]string, len(arr_1)*len(arr_2))
	for i := 0; i < len(arr_1); i++ {
		for j := 0; j < len(arr_2); j++ {
			idx := i*len(arr_2) + j
			res[idx] = arr_1[i] + arr_2[j]
		}
	}

	return res
}

func Search(s string) []string {
	if len(s) == 1 {
		num, _ := strconv.Atoi(s)
		return letterMap[num]
	}

	return Combine(Search(s[:1]), Search(s[1:]))
}

func main() {
	fmt.Println(letterCombinations("123"))
}
