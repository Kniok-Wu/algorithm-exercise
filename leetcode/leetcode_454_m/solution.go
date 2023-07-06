package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	record := make(map[int]int)
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			record[nums1[i]+nums2[j]] += 1
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			count += record[-(nums3[i] + nums4[j])]
		}
	}
	return count
}
