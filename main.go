package main

import (
	"fmt"
)

func main() {
	var price float32 = 89.12
	fmt.Println("price =", price)
	var num1 float32 = -0.00089
	var num2 float64 = -7809666.89
	fmt.Println("num1 =", num1, "num2 =", num2)
	// 尾數部分可能丟失，造成精度損失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 =", num3, "num4 =", num4)
}
