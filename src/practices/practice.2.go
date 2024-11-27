package main

import "fmt"

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func main() {
	fmt.Println(f())
}