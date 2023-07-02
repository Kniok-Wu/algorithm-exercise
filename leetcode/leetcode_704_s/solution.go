package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for end >= start {
		mid := int((end + start + 1) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func main() {
	//fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
