// 28. 实现 strStr()
// 实现 strStr() 函数。

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

// 说明：

// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

// 示例 1：

// 输入：haystack = "hello", needle = "ll"
// 输出：2
// 示例 2：

// 输入：haystack = "aaaaa", needle = "bba"
// 输出：-1
// 示例 3：

// 输入：haystack = "", needle = ""
// 输出：0

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	} else if len(needle) > len(haystack) {
		return -1
	} else {
		j := 0
		for i := 0; i < len(haystack)-len(needle)+1; i++ {
			if haystack[i] == needle[0] {
				for j = 0; j < len(needle); j++ {
					if haystack[i+j] != needle[j] {
						break
					}
				}
				if j == len(needle) {
					return i
				}
			}
		}
		return -1
	}
}