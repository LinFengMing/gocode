package main

import (
	"fmt"
)

func main() {
	var x int = 4
	var y int = 1
	if x > 2 {
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("atguigu")
	} else {
		fmt.Println("x is =", x)
	}
	var z int = 4
	if z > 2 {
		fmt.Println("ok")
	} else {
		fmt.Println("hello")
	}
}
