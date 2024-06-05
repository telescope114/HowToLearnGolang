package main
import (
  "fmt"
)

func main () {
	type Person struct {
		name string
		age int
	}
	person  := Person{name: "Alice", age: 20}
	fmt.Println(person)
	fmt.Println(person.name)
}
