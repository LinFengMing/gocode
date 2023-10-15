package main

import (
	"fmt"
)

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("result =", result)
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	result2 := a(10, 20)
	fmt.Println("result2 =", result2)
	result3 := Fun1(4, 9)
	fmt.Println("result3 =", result3)
}
