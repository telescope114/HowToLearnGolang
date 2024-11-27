package main

func romanToInt(s string) int {
	var romanToIntMap map[byte]int = map[byte]int{
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
	for i := 0; i < n; i++ {
		// 获取当前字符对应的整数值
		curr := romanToIntMap[s[i]]

		// 检查下一个字符（如果存在）是否比当前字符大
		if i < n-1 && romanToIntMap[s[i+1]] > curr {
			// 如果是特殊情况，比如 "IV" 或 "IX"，减去当前值
			total -= curr
		} else {
			// 正常累加当前值
			total += curr
		}
	}
	return total
}

func main() {
	// fmt.Println("hello world")
	romanToInt("LVIII")
}
