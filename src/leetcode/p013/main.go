package main

import "fmt"

func romanToInt(s string) int {
	// 使用数组替代map,性能更好
	romanToIntArr := [128]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	if n == 0 {
		return 0
	}

	total := 0
	prev := 0
	// 从右向左遍历,避免额外的判断
	for i := n - 1; i >= 0; i-- {
		curr := romanToIntArr[s[i]]
		if curr >= prev {
			total += curr
		} else {
			total -= curr
		}
		prev = curr
	}
	return total
}

func main() {
	// fmt.Println("hello world")
	fmt.Println(romanToInt("LVIII"))
}
