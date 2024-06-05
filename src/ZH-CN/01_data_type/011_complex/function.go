package main
import ("fmt")

// 函数声明
func base () {
	fmt.Println("this is base function")
}

func add (num1 int, num2 int) (int) {
	return num1 + num2
}

func main () {
	// 函数调用
	base()

	// 获取有返回值的函数
	sum := add(13, 22)
	fmt.Println(sum)
}