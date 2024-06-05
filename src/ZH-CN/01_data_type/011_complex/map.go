package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// 基础使用
    myMap := make(map[string]string)
    myMap["robot"] = "🤖"
    myMap["clown"] = "🤡"
    myMap["mushroom"] = "🍄"
    fmt.Println(myMap)
	if myMap["mushroom"] != "" {
		fmt.Println(myMap["mushroom"])
	} else {
		fmt.Println("蘑菇未被声明")
	}
}