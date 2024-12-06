package main

import "fmt"

func hasPrefix(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if str[i] != prefix[i] {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	prefix := strs[0]

	// 遍历所有字符串
	for i := 1; i < len(strs); i++ {
		// 当prefix不为空且当前字符串不是以prefix开头时
		for len(prefix) > 0 && !hasPrefix(strs[i], prefix) {
			// 移除最后一个字符重试
			prefix = prefix[:len(prefix)-1]
		}

		// 如果prefix已经为空,说明没有公共前缀
		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"internet", "interesting", "interview"}))
}
