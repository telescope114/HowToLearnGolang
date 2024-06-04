package main;
import (
	"fmt"
);

func main () {
	var P *int = nil;	// 声明 P 为int类型的指针
	var p int = 32;
	P = &p;				// 将 指针P 
	/* 
	p := 32;
	P := &p;
	 */
	fmt.Println(P);		// 单独打印只会打印出指针指向地址
	fmt.Println(*P);	// 添加*可以打印指针指向的值

	type Person struct {
		name string;
		age int64;
	}

	person := Person{name: "李华", age: 22};
	person_pointers := &person;
	fmt.Println(person_pointers);	// 这个不会打印指针地址，只会打印到对应的结构体本身
	fmt.Printf("%p", person_pointers);	// 这个会打印结构体对应的地址
}
