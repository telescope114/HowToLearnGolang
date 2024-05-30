package main;
import (
	"fmt"
);

func main () {
	var P *int = nil; // 声明P 为int类型的指针
	var p int = 32;
	P = &p;
	/* 
	p := 32;
	P := &p;
	 */
	fmt.Println(p);
	fmt.Println(P);
	fmt.Println(*P);
}
