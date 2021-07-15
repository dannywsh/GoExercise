func insertSort(nums []int) {
	var j int
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		for j = i - 1; j >= 0; j-- {
			if tmp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = tmp
				break
			}
		}
		if j < 0 {
			nums[j+1] = tmp
		}
	}
}
