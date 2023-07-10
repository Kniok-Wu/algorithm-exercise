package KMP

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	str := "BBC ABCDAB ABCDABCDABDE"
	subStr := "ABCDABD"
	idx := KMP(str, subStr)
	fmt.Println(str[idx : idx+len(subStr)])
}
