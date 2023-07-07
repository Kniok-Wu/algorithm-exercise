package Sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{2, 3}
	QuickSort(nums)
	fmt.Println(nums)
}
