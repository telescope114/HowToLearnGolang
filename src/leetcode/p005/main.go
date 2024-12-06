// go:build 005
package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) <= 2 {
		return s
	}
	return ""
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
