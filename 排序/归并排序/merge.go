//递归法
func sortArray(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, begin int, end int) {
	if begin < end {
		mid := begin + (end-begin)/2
		mergeSort(nums, begin, mid)
		mergeSort(nums, mid+1, end)
		merge(nums, begin, mid, end)
	}
}
func merge(nums []int, begin int, mid int, end int) {
	out := make([]int, 0)
	arr1Start := begin
	arr2Start := mid + 1
	for arr1Start <= mid && arr2Start <= end {
		if nums[arr1Start] < nums[arr2Start] {
			out = append(out, nums[arr1Start])
			arr1Start++
		} else {
			out = append(out, nums[arr2Start])
			arr2Start++
		}
	}
	if arr1Start > mid {
		out = append(out, nums[arr2Start:end+1]...)
	} else if arr2Start > end {
		out = append(out, nums[arr1Start:mid+1]...)
	}
	for _, v := range out {
		nums[begin] = v
		begin++
	}
}