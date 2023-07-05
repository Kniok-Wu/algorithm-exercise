package main

//func intersection(nums1 []int, nums2 []int) []int {
//	record := make(map[int]int)
//
//	for i := 0; i < len(nums1); i++ {
//		if _, ok := record[nums1[i]]; !ok {
//			record[nums1[i]] = 1
//		}
//	}
//
//	for i := 0; i < len(nums2); i++ {
//		if _, ok := record[nums2[i]]; ok {
//			record[nums2[i]] += 1
//		}
//	}
//	result := make([]int, 0)
//
//	for k, v := range record {
//		if v > 1 {
//			result = append(result, k)
//		}
//	}
//
//	return result
//}

func intersection(nums1 []int, nums2 []int) []int {
	record := make([]int, 1001)

	for i := 0; i < len(nums1); i++ {
		if record[nums1[i]] == 0 {
			record[nums1[i]] = 1
		}
	}

	for i := 0; i < len(nums2); i++ {
		if record[nums2[i]] == 1 {
			record[nums2[i]] = 2
		}
	}
	result := make([]int, 0)

	for k, v := range record {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}
