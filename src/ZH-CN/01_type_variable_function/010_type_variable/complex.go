/**
 复杂类型如下：
 * 指针
 * 数组
 * 结构体
 * 函数
 * 管道
 * 切片
 * 接口
 * map
 */

package main

import "fmt"


func main() {
	// Person 结构体类型声明
	type Person struct {
		name string
		age  int
	}
	// 创建一个 Person 类型的变量
	var P  = nil
	p := Person{name: "Alice", age: 20}
	P = &P
	fmt.Println(p) // 输出 "{Alice 20}"
	fmt.Println(P)

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
