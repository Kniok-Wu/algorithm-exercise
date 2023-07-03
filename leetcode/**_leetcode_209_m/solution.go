package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	result := l + 1
	i, k := 0, 0
	sum := 0

	for k < l {
		sum += nums[k]
		for sum >= target {
			if result > k-i+1 {
				result = k - i + 1
			}
			sum -= nums[i]
			i += 1
		}
		k += 1
	}
	if result == l+1 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}
