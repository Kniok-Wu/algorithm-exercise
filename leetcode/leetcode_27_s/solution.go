package main

import "fmt"

func removeElement(nums []int, val int) int {
	fast := 0
	slow := 0
	length := len(nums)

	for fast <= len(nums)-1 {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow += 1
		} else {
			length -= 1
		}
		fast += 1
	}

	return length
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	res := removeElement(nums, 2)
	fmt.Println(res, nums[:res])
}
