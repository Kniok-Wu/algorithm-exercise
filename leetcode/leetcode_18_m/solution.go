package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			base := nums[i] + nums[j]
			h := j + 1
			k := len(nums) - 1

			for h < k {
				if nums[h]+nums[k]+base > target {
					k -= 1
				} else if nums[h]+nums[k]+base < target {
					h += 1
				} else {
					result = append(result, []int{nums[i], nums[j], nums[h], nums[k]})
					for h+1 < k && nums[h+1] == nums[h] {
						h += 1
					}
					for k-1 > h && nums[k-1] == nums[k] {
						k -= 1
					}
					h += 1
					k -= 1
				}
			}

			for j+1 < len(nums)-2 && nums[j+1] == nums[j] {
				j += 1
			}
		}
		for i+1 < len(nums)-3 && nums[i+1] == nums[i] {
			i += 1
		}
	}

	return result
}

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
