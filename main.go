package main

import (
	"fmt"
)

func main() {
	// & 和 * 的使用
	a := 100
	fmt.Println("a 的地址是", &a)
	var ptr *int = &a
	fmt.Println("ptr 指向的值是", *ptr)
	var n int
	var i int = 10
	var j int = 12
	// golang 沒有三元運算符
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n =", n)
	// 取兩個數最大值
	var n1 int = 10
	var n2 int = 10
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max =", max)
	// 取三個數最大值
	var n3 = 45
	if n3 > max {
		max = n3
	}
	fmt.Println("max =", max)
}
