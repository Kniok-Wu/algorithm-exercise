package Sorts

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[0]

	left, right := 0, len(nums)-1

	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}

		if left < right {
			nums[left] = nums[right]
			left += 1
		}

		for left < right && nums[left] < pivot {
			left++
		}

		if left < right {
			nums[right] = nums[left]
			right -= 1
		}
	}

	nums[right] = pivot
	QuickSort(nums[:right])
	QuickSort(nums[right+1:])
}
