func sortArray(array1 []int) (res []int) {
	res = array1
	increament := 4
	for increament > 1 {
		increament /= 2
		var j int
		var tmp int
		for i := 0; i < increament; i++ {
			for s := i + increament; s < len(res); s += increament {
				tmp = res[s]
				for j = s - increament; j > -1; j -= increament {
					if tmp > res[j] {
						res[j+increament] = res[j]
					} else {
						res[j+increament] = tmp
						break
					}
				}
				res[j+increament] = tmp
			}

		}
	}
	return
}