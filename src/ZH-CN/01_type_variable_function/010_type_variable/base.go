package main

import (
	"fmt"
	"reflect"
)

func main() {
	//#region

	// 基础类型
	// 布尔型
	var b_1 bool = false
	b_2 := true // 这种声明会自动判断b_2的类型
	fmt.Println(reflect.TypeOf(b_1))
	fmt.Printf("%T\n", b_2)
	//#endregion

	//#region

	// 数字类型
	// 整型
	var i_1 int = 333 // 有符号整型，受限于当前操作系统位数，32位对应int32，64位对应int64
	i_2 := 444
	var i_3 uint = 123 // 无符号整型，只能表示非负数，受限于当前操作系统位数，32位对应uint32，64位对应uint64
	fmt.Println(i_1 + i_2)
	fmt.Println(reflect.TypeOf(i_1))
	fmt.Println(reflect.TypeOf(i_3))
	//fmt.Println(i_1 + i_3) // 类型不对无法运算，必须类型一致
	//i_1 = false // 已经设置好的类型变量无法被覆盖
	fmt.Println(uint(i_1) + i_3) // 可以通过类型转换进行运算

	// 浮点型
	var (
		f_1 float64 = 45.00
		f_2 float32 = 33.23
	)
	f_3 := 12.33
	fmt.Println(f_1 + f_3)
	fmt.Println(reflect.TypeOf(f_1 + f_3))
	fmt.Println(reflect.TypeOf(f_2))
	fmt.Println(f_2 + float32(f_1)) // 不建议大范围转小范围
	//#endregion

	//#region

	// 字符串型
	var str_1 string = "a123"
	fmt.Println(reflect.TypeOf(str_1))
	str_2 := 'a' // 注意：字符串只有双引号，否则会自动转化成字符码(int32)
	fmt.Println(reflect.TypeOf(str_2))
	var str_3 string = "bbbbb"
	fmt.Println(str_1 + str_3)
	//#endregion

	//#region

	// 空值 nil
	// 一般无法直接赋值给基础类型
	var p *int = nil // 这个*int代表指针变量
	var p1 int = 33
	p = &p1
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p)) // 类型还是初始声明的类型
	//#endregion
}
