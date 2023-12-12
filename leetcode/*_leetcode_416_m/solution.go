/**
 * @Time: 2023/12/12 16:29
 * @Author: kniokwu@gmail.com
 * @File: solution.go
 * @Software: GoLand Algorithm
 * @Description: $
 */

package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0

	for _, val := range nums {
		sum += val
	}
	capacity := int(sum / 2)
	if capacity*2 != sum {
		return false
	}

	dp := make([]int, capacity+1)
	for i := 0; i <= capacity; i++ {
		if i >= nums[0] {
			dp[i] = nums[0]
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := capacity; j >= 0; j-- {
			if j < nums[i] {
				break
			}

			if j >= nums[i] {
				curValue := nums[i] + dp[j-nums[i]]
				if curValue == capacity {
					return true
				}

				if curValue > dp[j] {
					dp[j] = curValue
				}
			}
		}
	}

	return false
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
