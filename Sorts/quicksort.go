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

func quickSort(nums []int, left int, right int) {
	if right-left < 1 {
		return
	}

	pivot, index := left, left+1

	for i := left; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index += 1
		}
	}

	nums[index-1], nums[pivot] = nums[pivot], nums[index-1]
	quickSort(nums, left, index-1)
	quickSort(nums, index+1, right)
}
