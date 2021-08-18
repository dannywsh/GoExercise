package main

import (
	"fmt"
)

func main() {
	a := printNumbers(2)
	fmt.Println(a)
}
func printNumbers(n int) []string {
	res := make([]string, 0)
	for digit := 1; digit <= n; digit++ {
		for j := '1'; j <= '9'; j++ {
			nums := make([]byte, digit)
			nums[0] = byte(j)
			dfs(1, nums, digit, &res)
		}
	}
	return res
}
func dfs(dep int, nums []byte, digit int, res *[]string) {
	if dep == digit {
		*res = append(*res, string(nums))
		return
	}

	for i := '0'; i <= '9'; i++ {
		nums[dep] = byte(i)
		dfs(dep+1, nums, digit, res)
	}
}
