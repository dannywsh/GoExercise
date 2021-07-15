func sortArray(array1 []int) (res []int) {
	res = array1
	var tmp int
	for i := 0; i < len(res)-1; i++ {
		tmp = i
		for j := i + 1; j < len(res); j++ {
			if res[tmp] < res[j] {
				tmp = j
			}
		}
		if i != tmp {
			res[i], res[tmp] = res[tmp], res[i]
		}
	}
	return
}