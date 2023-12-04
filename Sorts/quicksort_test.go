package Sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{7, 10, 2, 5, 8, 4, 3, 9, 15}
	nums = heapSort(nums)
	fmt.Println(nums)
}
