package main

import (
	"fmt"
)

func main() {
	// 如果運算的數都是整數，除法只會保留整數部分
	fmt.Println(10 / 4)
	// 要保留小數部分，需要有浮點數加入運算
	fmt.Println(10.0 / 4)
	// a % b = a - a / b * b
	fmt.Println("10 % 3 =", 10%3)
	fmt.Println("-10 % 3 =", -10%3)
	fmt.Println("10 % -3 =", 10%-3)
	fmt.Println("-10 % -3 =", -10%-3)
	// ++ and --
	var i int = 10
	i++
	fmt.Println("i =", i)
	i--
	fmt.Println("i =", i)
}
