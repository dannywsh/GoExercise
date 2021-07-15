//挖坑法
func sortArray(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}
func quickSort(nums []int, begin int, end int) {
	if begin < end {
		index := partition(nums, begin, end)
		quickSort(nums, begin, index-1)
		quickSort(nums, index+1, end)
	}
}
func partition(nums []int, low int, high int) int {
	pivot := nums[low]
	for low < high {
		for ; low < high && nums[high] > pivot; high-- {
		}
		if low < high {
			nums[low] = nums[high]
		}
		for ; low < high && nums[low] <= pivot; low++ {
		}
		if low < high {
			nums[high] = nums[low]
		}
	}
	nums[low] = pivot
	return low
}

//交换法
func sortArray(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}
func quickSort(nums []int, begin int, end int) {
	if begin < end {
		index := partition(nums, begin, end)
		quickSort(nums, begin, index-1)
		quickSort(nums, index+1, end)
	}
}
func partition(nums []int, low int, high int) int {
	pivot := low
	for low < high {
		for ; low < high && nums[high] > nums[pivot]; high-- {
		}
		for ; low < high && nums[low] <= nums[pivot]; low++ {
		}
		if low < high {
			nums[low], nums[high] = nums[high], nums[low]
		}
	}
	nums[low], nums[pivot] = nums[pivot], nums[low]
	return low
}