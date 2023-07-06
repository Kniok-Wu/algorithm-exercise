package main

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok2 := record[target-nums[i]]; ok2 {
			return []int{idx, i}
		}

		record[nums[i]] = i
	}

	return nil
}
