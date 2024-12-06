package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	base := strconv.FormatInt(int64(x), 10)
	runes := []rune(base)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return base == string(runes)
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
