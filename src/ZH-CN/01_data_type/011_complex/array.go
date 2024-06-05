package main;

import "fmt";

func main() {
    // 示例 1：简单定义和使用数组
    arr1 := [3]int{10, 20, 30};
    fmt.Println(arr1);

    // 示例 2：遍历数组
    for i, v := range arr1 {
        fmt.Printf("index is %d, value is %d\n", i, v);
    }

    // 示例 3：定义未初始化的数组
    var arr2 [5]int;
    arr2[0] = 5;
    arr2[1] = 15;
    fmt.Println(arr2);

	// 示例 4：使用结构体一起完成
	type Person struct {
		name string
		age int
	};
	person := [5]Person{
		{ name: "张三", age: 20 },
		{ name: "李四", age: 21 },
		{ name: "王五", age: 22 },
		{ name: "赵六", age: 23 },
		{ name: "田七", age: 24 },
	};
	p_pointer := &person;
	fmt.Println(person);
	fmt.Printf("%p", p_pointer);
}