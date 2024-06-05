package main;

import (
	"fmt"
);

type Person struct {
	Name string
	Age int
}

func main () {
	numbers := make([]int, 3, 5);
	numbers[0] = 33;
	fmt.Println(numbers);
	fmt.Println(len(numbers));
	// 切片扩容
	numbers = append(numbers, 44, 55);
	fmt.Println(numbers);
	fmt.Println(len(numbers));


	// 结构体类型的数组也可以
	person := make([]Person, 10);
	for index := range person {
		person[index] = Person{Name: "张三", Age: index};
	}
	fmt.Println(person, len(person));
	// 同样这个可以通过append进行扩容
	person = append(person, Person{Name: "李四", Age: 222});
	fmt.Println(person, len(person));
}
