package main;

import (
	"fmt"
);

func lengthOfLongestSubstring(s string) int {
	maxLen := 0;

	for i := 0; i < len(s); i ++ {
		for j := i + 1; j < len(s); j ++ {
			var checked bool = true;
			for h := i; h < j; h ++ {
				if (s[j] == s[h]) {
					checked = false;
					break;
				}
			} 
			if checked {
				if (j - i) > maxLen {
					maxLen = j - i
				}
			} else {
				break
			}
		}
	}

	return maxLen + 1;
}

func main () {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"));
}