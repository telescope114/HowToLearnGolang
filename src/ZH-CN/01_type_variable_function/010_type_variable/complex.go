package main

import "fmt"

// Person 类型声明
type Person struct {
	name string
	age  int
}

func main() {
	// 创建一个 Person 类型的变量
	p := Person{name: "Alice", age: 20}
	fmt.Println(p) // 输出 "{Alice 20}"

	// 创建一个切片类型的变量
	s := []int{1, 2, 3}
	fmt.Println(s[0]) // 输出 "[1 2 3]"

	// 创建一个映射类型的变量
	m := map[int]Person{
		0: {name: "Alice0", age: 120},
		1: {name: "Alice1", age: 121},
		2: {name: "Alice2", age: 122},
		3: {name: "Alice3", age: 123},
	}
	fmt.Println(m[0]) // 输出 "map[apple:1 banana:2]"
	fmt.Println(m)
}
