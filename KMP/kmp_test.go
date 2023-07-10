package KMP

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	fmt.Println(KMP("BBC ABCDAB ABCDABCDABDE", "ABCDABD"))
}
