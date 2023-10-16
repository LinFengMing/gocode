package main

import (
	"fmt"
)

var age int = 50
var Name string = "Jack"

func test() {
	age := 10
	Name := "Tom"
	fmt.Println("age =", age)
	fmt.Println("Name =", Name)
}

func main() {
	test()
	fmt.Println("age =", age)
	fmt.Println("Name =", Name)
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
	fmt.Println("i =", i)
}
