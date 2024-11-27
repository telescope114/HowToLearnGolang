package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int64
}

func main() {
	var persons = []Person{
		{name: "张三", age: 20},
		{name: "李四", age: 21},
		{name: "王五", age: 22},
		{name: "赵六", age: 23},
		{name: "田七", age: 24},
	}

	// 可以用 range 来获取 数组/切片 的序数和值
	for i, v := range persons {
		fmt.Printf("第%d个人是%s，年龄%d。\n", i, v.name, v.age)
	}

	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i])
	}
}
