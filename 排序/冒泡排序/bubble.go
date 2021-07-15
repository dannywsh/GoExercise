//最简单的冒泡
func sortArray(array1 []int) (res []int) {
	res = array1
	for i := 0; i < len(res)-1; i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] < res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return
}

//考虑有序后的冒泡
func sortArray(array1 []int) (res []int) {
	res = array1
	flag := true //设置标志位
	for i := 0; i < len(res)-1 && flag; i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] < res[j+1] {
				flag = false
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return
}